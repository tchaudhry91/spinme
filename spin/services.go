/*
Copyright © 2019 Tanmay Chaudhry <tanmay.chaudhry@gmail.com>

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
package spin

import (
	"errors"
	"strings"
)

// SpinnerFrom returns the required SpinnerFunc from the service string
func SpinnerFrom(service string) (SpinnerFunc, error) {
	switch s := strings.ToLower(service); s {
	case "mongo":
		return Mongo, nil
	case "postgres":
		return Postgres, nil
	case "mysql":
		return MySQL, nil
	case "redis":
		return Redis, nil
	default:
		return Generic, errors.New("Failed to find given service")
	}
}
