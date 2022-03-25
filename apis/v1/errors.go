// Copyright 2021-2022 The phy-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import "fmt"

func IsError400(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails400)
	return ok
}

func IsError401(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails401)
	return ok
}

func IsError404(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails404)
	return ok
}

func IsError409(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails409)
	return ok
}

func IsError429(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails429)
	return ok
}

func IsError503(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ProblemDetails503)
	return ok
}

func (e ProblemDetails400) Error() string {
	return fmt.Sprintf("status: %d, title: %v, detail: %v", e.Status, e.Title, e.Detail)
}

func (e ProblemDetails401) Error() string {
	return fmt.Sprintf("Code: %v, status: %s, msg: %v", e.ErrorCode, e.Status, e.ErrorMsg)
}

func (e ProblemDetails404) Error() string {
	return fmt.Sprintf("status: %d, title: %v, detail: %v", e.Status, e.Title, e.Detail)
}

func (e ProblemDetails409) Error() string {
	return fmt.Sprintf("status: %d, title: %v, detail: %v", e.Status, e.Title, e.Detail)
}

func (e ProblemDetails429) Error() string {
	return fmt.Sprintf("status: %d, title: %v, detail: %v", e.Status, e.Title, e.Detail)
}

func (e ProblemDetails503) Error() string {
	return fmt.Sprintf("status: %d, title: %v, detail: %v", e.Status, e.Title, e.Detail)
}
