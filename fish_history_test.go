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

package fht_test

import (
	"regexp"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)
import "livingit.de/code/fht"

// helper functions

func getEchoEntry() *fht.FishHistory {
	return &fht.FishHistory{
		Command: "echo hello",
		When:    0,
	}
}

func getPathEntry() *fht.FishHistory {
	return &fht.FishHistory{
		Command: "mv /home/user/directory/* /home/user/otherdirectory/",
		When:    0,
		Paths:   []string{"/home/user/directory", "/home/user/otherdirectory"},
	}
}

// Tests

// Tests for stringer interface

type stringerTest struct {
	Name     string
	History  *fht.FishHistory
	Expected string
}

var stringerTests = []stringerTest{
	{
		Name:    "echo",
		History: getEchoEntry(),
		Expected: `- cmd: echo hello
  when: 0`,
	},
	{
		Name:    "paths",
		History: getPathEntry(),
		Expected: `- cmd: mv /home/user/directory/* /home/user/otherdirectory/
  when: 0
  paths:
    - /home/user/directory
    - /home/user/otherdirectory`,
	},
}

func TestStringer(t *testing.T) {
	for _, test := range stringerTests {
		t.Run(test.Name, func(t *testing.T) {
			result := test.History.String()
			if result != test.Expected {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(test.Expected, result, true)

				t.Logf("Expected:\n  [%s]\nGot:\n  [%s]\nDiff:\n%s", test.Expected, result, dmp.DiffPrettyText(diffs))
				t.Fail()
			}
		})
	}
}

// Tests for Matches

type matchesTest struct {
	Name       string
	History    *fht.FishHistory
	Expression string
	Expected   bool
}

var matchesTests = []matchesTest{
	{
		Expected:   true,
		History:    getEchoEntry(),
		Name:       "echo - match",
		Expression: "echo",
	},
	{
		Expected:   false,
		History:    getPathEntry(),
		Name:       "echo - no match",
		Expression: "echo",
	},
	{
		Expected:   true,
		History:    getPathEntry(),
		Name:       "path - match",
		Expression: ".*directory",
	},
}

func TestMatches(t *testing.T) {
	for _, test := range matchesTests {
		t.Run(test.Name, func(t *testing.T) {
			expression, err := regexp.Compile(test.Expression)
			if err != nil {
				t.Logf("expression invalid: [%s]", err)
			}
			result, err := test.History.Matches(expression)
			if result != test.Expected {
				t.Fail()
			}
		})
	}
}
