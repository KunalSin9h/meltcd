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

package meltcd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/meltred/meltcd/internal/core"
	"github.com/meltred/meltcd/util"
	"github.com/spf13/cobra"
)

func LoginUser(cmd *cobra.Command, _ []string) error {
	showToken, _ := cmd.Flags().GetBool("show-token")
	if showToken {
		fmt.Println(core.GetAccessToken())
		return nil
	}

	username, err := stringPrompt("Username")
	if err != nil {
		return err
	}

	passwordBytes, err := util.GetUsersPassword("Password: ", true, os.Stdin, os.Stderr)
	if err != nil {
		return err
	}

	password := string(passwordBytes)
	if err := checkEmpty(password); err != nil {
		return err
	}

	reqBody := map[string]string{
		"username": username,
		"password": password,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(reqBody); err != nil {
		return err
	}

	res, err := http.Post(fmt.Sprintf("%s/api/login", util.GetServer()), "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resData := string(resDataBytes)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("login failed: %s", resData)
		return err
	}

	if err := core.StoreAccessToken(resData); err != nil {
		return err
	}

	color.HiGreen("Login Success")
	return nil
}

func stringPrompt(label string) (string, error) {
	r := bufio.NewReader(os.Stdin)

	fmt.Fprint(os.Stderr, label+": ")

	s, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	s = strings.TrimSpace(s)

	if err := checkEmpty(s); err != nil {
		return "", err
	}

	return s, nil
}

func checkEmpty(s string) error {
	if len(s) == 0 {
		err := fmt.Errorf("input can't be empty")
		return err
	}
	return nil
}
