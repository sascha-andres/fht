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
