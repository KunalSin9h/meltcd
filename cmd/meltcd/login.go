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
	"encoding/base64"
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

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/login", util.GetServer()), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodeToBasicAuth(username, password)))

	client := &http.Client{}

	res, err := client.Do(req)
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

// A function used to convert username, password to basic auth token
func encodeToBasicAuth(username, password string) string {
	userPass := fmt.Sprintf("%s:%s", username, password)

	return base64.StdEncoding.EncodeToString([]byte(userPass))
}
