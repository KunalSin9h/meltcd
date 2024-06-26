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

package version

// Version will be altered at CI step when new version release
// using ldflags we can change the version at compile time
// go build -ldflags "-X github.com/kunalsin9h/meltcd/version.Version=0.3.1" main.go
var Version string = "dev" // nolint:all
