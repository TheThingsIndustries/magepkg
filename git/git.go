// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package git

import (
	"fmt"
	"strings"

	"github.com/magefile/mage/sh"
)

// ChangedFiles returns the files that changed between the given commits.
func ChangedFiles(previousSHA, currentSHA string) ([]string, error) {
	out, err := sh.Output("git", "diff", "--name-only", fmt.Sprintf("%s...%s", previousSHA, currentSHA))
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

// StagedFiles returns the files that are staged by git.
func StagedFiles() ([]string, error) {
	out, err := sh.Output("git", "diff", "--staged", "--name-only", "--diff-filter=d")
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

// UnstagedFiles returns the files that are changed, but not staged by git.
func UnstagedFiles() ([]string, error) {
	out, err := sh.Output("git", "diff", "--name-only", "--diff-filter=d")
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

// Info returns Git info.
func Info() (commit, branch, tag string, err error) {
	commit, err = sh.Output("git", "rev-parse", "HEAD")
	if err != nil {
		return
	}
	branch, err = sh.Output("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return
	}
	tag, err = sh.Output("git", "describe", "--abbrev=0", "--tags")
	if err != nil {
		err = nil // ignore error
	}
	return
}
