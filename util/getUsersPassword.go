package util

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/term"
)

type FieldReader interface {
	io.Reader
	Fd() uintptr
}

var readChunk = func(r io.Reader) (byte, error) {
	buf := make([]byte, 1)
	if n, err := r.Read(buf); n == 0 || err != nil {
		if err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	return buf[0], nil
}

var (
	maxLength            = 512
	ErrInterrupted       = errors.New("interrupted")
	ErrMaxLengthExceeded = fmt.Errorf("maximum byte limit (%v) exceeded", maxLength)
	chunk                = readChunk
)

// GetUsersPassword returns the input read from term.
// If prompt is not empty, it will be output as a prompt to the user
// If masked is true, typing will be matched by asterisks on the screen.
// Otherwise, typing will echo nothing.

func GetUsersPassword(msg string, masked bool, r FieldReader, w io.Writer) ([]byte, error) {
	var err error
	var p, bs, ms []byte
	if masked {
		bs = []byte("\b \b")
		ms = []byte("*")
	}

	if term.IsTerminal(int(r.Fd())) {
		if oldState, err := term.MakeRaw(int(r.Fd())); err != nil {
			return p, err
		} else { // nolint
			defer func() {
				err := term.Restore(int(r.Fd()), oldState)
				if err != nil {
					return
				}
				_, err = fmt.Fprintln(w)
				if err != nil {
					return
				}
			}()
		}
	}

	if msg != "" {
		_, err = fmt.Fprint(w, msg)
		if err != nil {
			return nil, err
		}
	}

	// Track total bytes read, not just bytes in the password.  This ensures any
	// errors that might flood the console with nil or -1 bytes infinitely are
	// capped.
	var count int
	for count = 0; count <= maxLength; count++ {
		if v, e := chunk(r); e != nil {
			err = e
			break
		} else if v == 127 || v == 8 { // nolint
			if l := len(p); l > 0 {
				p = p[:l-1]
				_, err := fmt.Fprint(w, string(bs))
				if err != nil {
					return nil, err
				}
			}
		} else if v == 13 || v == 10 {
			break
		} else if v == 3 {
			err = ErrInterrupted
			break
		} else if v != 0 {
			p = append(p, v)
			_, err = fmt.Fprint(w, string(ms))
			if err != nil {
				return nil, err
			}
		}
	}
	if count > maxLength {
		err = ErrMaxLengthExceeded
	}
	return p, err
}
