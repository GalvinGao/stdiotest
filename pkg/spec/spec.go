package spec

import (
	"os/exec"

	"github.com/GalvinGao/stdiotest/pkg/configor"
)

type TestCase struct {
	Cmd *exec.Cmd

	ExitCode int

	Stdin  string
	Stdout string
}

func NewTestCaseFromConfig(cfg *configor.TestCase) *TestCase {
	return &TestCase{
		Cmd:      exec.Command(cfg.Cmd, cfg.Args...),
		ExitCode: cfg.ExitCode,
		Stdin:    cfg.Stdin,
		Stdout:   cfg.Stdout,
	}
}
