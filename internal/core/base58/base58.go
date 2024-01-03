/*
Copyright 2023 - PRESENT Meltred

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package base58

import (
	"crypto/rand"
	"errors"
)

// defaultAlphabet is the alphabet used for ID characters by default.
// BASE-58
var defaultAlphabet = []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

const (
	defaultSize = 21
)

// New generates secure URL-friendly unique ID.
// Accepts optional parameter - length of the ID to be generated (21 by default).
func New(l ...int) (string, error) {
	var size int
	switch {
	case len(l) == 0:
		size = defaultSize
	case len(l) == 1:
		size = l[0]
		if size < 0 {
			return "", errors.New("negative id length")
		}
	default:
		return "", errors.New("unexpected parameter")
	}
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = defaultAlphabet[bytes[i]&57]
	}
	return string(id[:size]), nil
}
