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

package spec

import (
	"testing"
)

func TestNormalizeFilePath(t *testing.T) {
	testCaseNeg := []string{
		".env",
		"file.txt",
		"~/.service.yaml",
	}

	for _, file := range testCaseNeg {
		res, err := normalizeFilePath(file)
		if err != nil {
			t.Error(err.Error())
		}
		// when we normalize filePath it become like
		// file.txt    /home/user/directory/file.txt
		// ~/file.txt /home/user/file.txt
		//
		// so the length should not equal to expanded result
		if len(res) == len(file) {
			t.Log("file does not expanded", file)
			t.Fail()
		}
	}

	homePath := "/home/user/secret/something"
	res, err := normalizeFilePath(homePath)
	if err != nil {
		t.Error(err.Error())
	}

	if res != homePath {
		t.Log("file expanded even if the path is home", "path", homePath)
		t.Fail()
	}
}
