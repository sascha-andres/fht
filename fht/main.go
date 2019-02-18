package main

import (
	"fmt"
	"io/ioutil"
	"livingit.de/code/fht/parser"
	"os"
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

	_, err = p.ParseString(string(data))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error parsing: %s\n", err)
		os.Exit(1)
	}
}
