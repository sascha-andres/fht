package main

import (
	"fmt"
	"io/ioutil"
	"livingit.de/code/fht/parser"
	"os"
	"regexp"
)

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
	delete := false
	expression := ""

	arguments := os.Args
	for _, argument := range arguments[1:] {
		switch argument {
		case "--help":
			showHelp = true
		case "--delete":
			delete = true
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

	if showHelp {
		fmt.Println("usage fht [--delete] <expression>")
		os.Exit(0)
	}

	r, err := regexp.Compile(expression)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error compiling regex: %s", err)
		os.Exit(1)
	}

	for _, entry := range history {
		isMatch, _ := entry.Matches(r)
		if isMatch && delete {
			continue
		}
		if !isMatch && !delete {
			continue
		}
		_, _ = fmt.Fprintln(os.Stdout, &entry)
	}
}
