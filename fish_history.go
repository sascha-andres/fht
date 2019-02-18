package fht

// FishHistory represents one command in fish history
type FishHistory struct {
	Command string   `yaml:"cmd"`
	paths   []string `yaml:"Paths"`
	when    int      `yaml:"When"`
}
