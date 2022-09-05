package runner

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/GalvinGao/stdiotest/pkg/ioprinter"
	"github.com/GalvinGao/stdiotest/pkg/spec"
)

var (
	ErrExitCodeMismatch = errors.New("exit code mismatch")
	ErrStdoutMismatch   = errors.New("stdout mismatch")
)

type Run struct {
	Spec *spec.TestCase

	Errors []error
}

func New(t *spec.TestCase) *Run {
	return &Run{
		Spec: t,
	}
}

func (r *Run) Start() {
	r.Errors = make([]error, 0)

	buf := bytes.NewBufferString(r.Spec.Stdin)

	pipe, err := r.Spec.Cmd.StdinPipe()
	if err != nil {
		log.Error().Err(err).Msg("failed to create stdin pipe")
		r.Errors = append(r.Errors, err)
		return
	}

	_, err = buf.WriteTo(pipe)
	if err != nil {
		log.Error().Err(err).Msg("failed to write to stdin pipe")
		r.Errors = append(r.Errors, err)
		return
	}

	err = pipe.Close()
	if err != nil {
		log.Error().Err(err).Msg("failed to close stdin pipe")
		r.Errors = append(r.Errors, err)
		return
	}

	// err = r.Spec.Cmd.Start()
	// if err != nil {
	// 	log.Error().Err(err).Msg("failed to start command")
	// 	r.Errors = append(r.Errors, err)
	// 	return
	// }

	output, err := r.Spec.Cmd.Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)

		if !ok {
			log.Error().Err(err).Msg("failed to get command output")
			r.Errors = append(r.Errors, err)
			return
		}

		if exitErr.ExitCode() != r.Spec.ExitCode {
			log.Error().Err(ErrExitCodeMismatch).Msgf("exit code mismatch: expected %d, got %d", r.Spec.ExitCode, exitErr.ExitCode())
			r.Errors = append(r.Errors, ErrExitCodeMismatch)
			return
		}

		if string(output) != r.Spec.Stdout {
			fmt.Println(ioprinter.Diff(r.Spec.Stdout, string(output)))
			e := errors.Wrapf(ErrStdoutMismatch, "got %s (expecting %s)", string(output), r.Spec.Stdout)
			r.Errors = append(r.Errors, e)
		}

	}
}
