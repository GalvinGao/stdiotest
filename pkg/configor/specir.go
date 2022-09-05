package configor

type Spec struct {
	TestCases []*TestCase `yaml:"test_cases"`
}

type TestCase struct {
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args"`

	ExitCode int `yaml:"exit_code"`

	Stdin  string `yaml:"stdin"`
	Stdout string `yaml:"stdout"`
}
