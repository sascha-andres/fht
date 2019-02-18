package main

import (
	"fmt"
	"io/ioutil"
	"livingit.de/code/fht/parser"
	"os"
	"regexp"
)

const versionNumber = "develop"

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
		_, _ = fmt.Fprintln(os.Stdout, versionNumber)
		os.Exit(0)
	}

	if showHelp {
		_, _ = fmt.Fprintln(os.Stdout, "fish shell history tool")
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
