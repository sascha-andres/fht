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

package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"livingit.de/code/fht"
)

// Parser is the fish history file parser
// There parser works on a few assumptions:
//   - a line starting with a dash (-) denotes a new entry
//   - all lines belonging to an entry are properly indented
//     (if not the dash for paths would trigger a new entry)
type Parser struct{}

// NewParser returns a parser
func NewParser() (*Parser, error) {
	return &Parser{}, nil
}

// ParseFile loads a history file into memory and passed over to ParseString
func (parser *Parser) ParseFile(file string) ([]fht.FishHistory, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return parser.ParseString(string(data))
}

// ParseString creates an array of FishHistory out of the passed data
func (parser *Parser) ParseString(data string) ([]fht.FishHistory, error) {
	if data == "" {
		return nil, errors.New("empty data provided")
	}
	var currentHistory *fht.FishHistory
	result := make([]fht.FishHistory, 0)
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		newline := scanner.Text()
		if strings.HasPrefix(newline, "- ") {
			if nil != currentHistory {
				result = append(result, *currentHistory)
			}
			currentHistory = &fht.FishHistory{}
			newline = strings.TrimPrefix(newline, "- ")
		}
		trimmed := strings.Trim(newline, " ")
		lineData := strings.SplitN(trimmed, ":", 2)
		if lineData[0] == "cmd" {
			currentHistory.Command = strings.Trim(lineData[1], " ")
			continue
		}
		if lineData[0] == "when" {
			err := getWhen(currentHistory, lineData[1])
			if err != nil {
				return nil, err
			}
			continue
		}
		if lineData[0] == "paths" {
			currentHistory.Paths = make([]string, 0)
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			currentHistory.Paths = append(currentHistory.Paths, strings.TrimPrefix(trimmed, "- "))
		}
	}
	if nil != currentHistory {
		result = append(result, *currentHistory)
	}
	return result, nil
}

func newEntry(newline string) *fht.FishHistory {
	return &fht.FishHistory{
		Command: strings.Trim(strings.SplitN(newline, ":", 2)[1], " "),
	}
}

func getWhen(currentHistory *fht.FishHistory, data string) (err error) {
	trimmed := strings.Trim(data, " ")
	converted, err := strconv.ParseUint(trimmed, 0, 0)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not convert when to uint: %s - %s", strings.Trim(data, " "), err))
	}
	currentHistory.When = converted
	return
}
