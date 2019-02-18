package fht

import (
	"errors"
	"fmt"
	"regexp"
)

// FishHistory represents one command in fish history
type FishHistory struct {
	Command string   `yaml:"cmd"`   // Command contains the executed command
	Paths   []string `yaml:"Paths"` // paths indicate which arguments were file paths, as a hint to the autosuggestion machinery
	When    int      `yaml:"When"`  // When is the time of history entry in unix time
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
