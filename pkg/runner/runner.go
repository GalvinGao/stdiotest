package runner

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Runner struct {
	Concurrency int
	Runs        []*Run
	Verbose     bool
}

func (r *Runner) check() error {
	if r.Concurrency < 1 {
		return errors.New("invalid concurrency")
	}

	if len(r.Runs) == 0 {
		return errors.New("no runs")
	}

	return nil
}

func (r *Runner) Start() error {
	if err := r.check(); err != nil {
		return err
	}

	var (
		wg      sync.WaitGroup
		limiter = make(chan struct{}, r.Concurrency)
	)
	for i, run := range r.Runs {
		limiter <- struct{}{}
		wg.Add(1)

		go func(run *Run, i int) {
			defer func() {
				<-limiter
				wg.Done()
			}()

			run.Start()

			if len(run.Errors) > 0 {
				log.Error().Errs("errors", run.Errors).Msgf("run %d failed", i)
			} else {
				log.Info().Msgf("run %d passed", i)
			}
		}(run, i)
	}

	wg.Wait()

	return nil
}
