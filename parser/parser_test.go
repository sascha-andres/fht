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

import "testing"

func TestParserSingle(t *testing.T) {
	p, _ := NewParser()
	result, err := p.ParseString(`- cmd: cat /alpha | go run main.go
  when: 100
  paths:
    - /alpha
    - main.go`)
	if err != nil {
		t.Logf("error parsing: %s", err)
		t.Fail()
	}
	if len(result) != 1 {
		t.Logf("expected an array of length 1 - received %d", len(result))
		t.Fail()
	}
	if result[0].When != 100 {
		t.Logf("expected a when value of 100 - received %d", result[0].When)
		t.Fail()
	}
	if result[0].Command != "cat /alpha | go run main.go" {
		t.Logf("expected a Command value of 'cat /alpha | go run main.go' - received [%s]", result[0].Command)
		t.Fail()
	}
	if len(result[0].Paths) != 2 {
		t.Logf("expected a an array length of 2 for Paths - received [%d]", len(result[0].Paths))
		t.Fail()
	}
	if result[0].Paths[0] != "/alpha" {
		t.Logf("expected a path value of '/alpha' - received [%s]", result[0].Paths[0])
		t.Fail()
	}
	if result[0].Paths[1] != "main.go" {
		t.Logf("expected a path value of 'main.go' - received [%s]", result[0].Paths[1])
		t.Fail()
	}
}

func TestParserMultiple(t *testing.T) {
	p, _ := NewParser()
	result, err := p.ParseString(`- cmd: cat /alpha | go run main.go
  when: 100
  paths:
    - /alpha
    - main.go
- cmd: cat /alpha | go run main.go
  when: 100
  paths:
    - /alpha
    - main.go`)
	if err != nil {
		t.Logf("error parsing: %s", err)
		t.Fail()
	}
	if len(result) != 2 {
		t.Logf("expected an array of length 2 - received %d", len(result))
		t.Fail()
	}
	if result[1].When != 100 {
		t.Logf("expected a when value of 100 - received %d", result[0].When)
		t.Fail()
	}
	if result[1].Command != "cat /alpha | go run main.go" {
		t.Logf("expected a Command value of 'cat /alpha | go run main.go' - received [%s]", result[0].Command)
		t.Fail()
	}
	if len(result[1].Paths) != 2 {
		t.Logf("expected a an array length of 2 for Paths - received [%d]", len(result[0].Paths))
		t.Fail()
	}
	if result[1].Paths[0] != "/alpha" {
		t.Logf("expected a path value of '/alpha' - received [%s]", result[0].Paths[0])
		t.Fail()
	}
	if result[1].Paths[1] != "main.go" {
		t.Logf("expected a path value of 'main.go' - received [%s]", result[0].Paths[1])
		t.Fail()
	}
}
