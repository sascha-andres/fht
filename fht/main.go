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

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"livingit.de/code/fht/parser"
)

const versionNumber = "develop"
const date = "develop"
const commit = "HEAD"

func printVersion() {
	_, _ = fmt.Fprintf(os.Stdout, "version:     %s\n", versionNumber)
	_, _ = fmt.Fprintf(os.Stdout, "compiled on: %s\n", date)
	_, _ = fmt.Fprintf(os.Stdout, "commit:      %s\n", commit)
}

func main() {
	stat, _ := os.Stdin.Stat()
	if !((stat.Mode() & os.ModeCharDevice) == 0) {
		_, _ = fmt.Fprintln(os.Stderr, "expecting fish history to be piped")
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error reading piped data: %s\n", err)
		os.Exit(1)
	}

	p, err := parser.NewParser()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error creating parser: %s\n", err)
		os.Exit(1)
	}

	history, err := p.ParseString(string(data))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error parsing: %s\n", err)
		os.Exit(1)
	}

	showHelp := false
	removeEntries := false
	expression := ""
	showVersion := false

	arguments := os.Args
	for _, argument := range arguments[1:] {
		switch argument {
		case "--help":
			showHelp = true
		case "--delete":
			removeEntries = true
		case "--version":
			showVersion = true
		default:
			if "" == expression {
				expression = argument
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "only one expression allowed")
				os.Exit(1)
			}
		}
	}

	if !showHelp && expression == "" {
		showHelp = true
	}

	if showVersion {
		printVersion()
		os.Exit(0)
	}

	if showHelp {
		_, _ = fmt.Fprintln(os.Stdout, "fish shell history tool")
		_, _ = fmt.Fprintln(os.Stdout, "")
		printVersion()
		_, _ = fmt.Fprintln(os.Stdout, "")
		_, _ = fmt.Fprintln(os.Stdout, "usage fht [--delete] [--help] <expression>")
		_, _ = fmt.Fprintln(os.Stdout, "  --delete  remove matched entries instead of limiting to")
		_, _ = fmt.Fprintln(os.Stdout, "  --version show version number")
		_, _ = fmt.Fprintln(os.Stdout, "  --help    show help")
		os.Exit(0)
	}

	r, err := regexp.Compile(expression)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error compiling regex: %s", err)
		os.Exit(1)
	}

	for _, entry := range history {
		isMatch, _ := entry.Matches(r)
		if isMatch && removeEntries {
			continue
		}
		if !isMatch && !removeEntries {
			continue
		}
		_, _ = fmt.Fprintln(os.Stdout, &entry)
	}
}
