/*
Copyright 2023 - PRESENT kunalsin9h

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

package repository

import (
	"errors"
)

func Update(url, image, username, password string) error {
	// eight url is empty or image is empty
	// so combining them will give the name
	repo, found := FindRepo(url + image)
	if !found {
		return errors.New("repository does not exists")
	}

	repo.saveCredential(username, password)
	return nil
}
