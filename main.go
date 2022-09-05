package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/urfave/cli/v2"

	"github.com/GalvinGao/stdiotest/pkg/configor"
	"github.com/GalvinGao/stdiotest/pkg/runner"
	"github.com/GalvinGao/stdiotest/pkg/spec"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	}).
		With().
		Timestamp().
		Logger()

	app := &cli.App{
		Name:        "stdiotest",
		Description: "stdiotest is a testing utility which tests the stdout output of a given program with specified stdin. It also supports running the tests in parallel to improve efficiency of testing.",
		Commands: []*cli.Command{
			{
				Name:        "run",
				Description: "Run the tests specified in the config file.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "parallel",
						Aliases: []string{"p"},
						Usage:   "number of tests to run in parallel",
						Value:   1,
					},

					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "verbose output",
						Value:   false,
					},
				},
				Action: func(c *cli.Context) error {
					conf := configor.Parse()
					p := c.Int("parallel")
					v := c.Bool("verbose")

					runs := make([]*runner.Run, 0, len(conf.TestCases))
					for _, t := range conf.TestCases {
						runs = append(runs, runner.New(spec.NewTestCaseFromConfig(t)))
					}

					r := runner.Runner{
						Concurrency: p,
						Runs:        runs,
						Verbose:     v,
					}
					return r.Start()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}
