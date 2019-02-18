// Copyright © 2019 Sascha Andres <sascha.andres@outlook.com>
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

package fht

import (
	"errors"
	"fmt"
	"regexp"
)

// FishHistory represents one command in fish history
type FishHistory struct {
	Command string   // Command contains the executed command
	Paths   []string // paths indicate which arguments were file paths, as a hint to the autosuggestion machinery
	When    uint64   // When is the time of history entry in unix time
}

// String implements the stringer interface
func (entry *FishHistory) String() string {
	pathSection := ""
	for _, val := range entry.Paths {
		if pathSection == "" {
			pathSection = "\n  paths:"
		}
		pathSection = fmt.Sprintf("%s\n    - %s", pathSection, val)
	}
	return fmt.Sprintf("- cmd: %s\n  when: %d%s", entry.Command, entry.When, pathSection)
}

// Matches returns if an entry matches a regular expression
func (entry *FishHistory) Matches(r *regexp.Regexp) (bool, error) {
	if nil == r {
		return false, errors.New("no regular expression provided")
	}
	return entry.matchesCommand(r) || entry.matchesPath(r), nil
}

// matchesCommand applies regex to the command
func (entry *FishHistory) matchesCommand(r *regexp.Regexp) bool {
	return r.MatchString(entry.Command)
}

// matchesPath applies regex to paths until found / not found
func (entry *FishHistory) matchesPath(r *regexp.Regexp) bool {
	for _, val := range entry.Paths {
		if r.MatchString(val) {
			return true
		}
	}
	return false
}
