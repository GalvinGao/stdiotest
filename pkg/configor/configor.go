package configor

import (
	"os"

	"github.com/go-yaml/yaml"
	"github.com/rs/zerolog/log"
)

func Parse() Spec {
	f, err := os.OpenFile("stdiotest.yaml", os.O_RDONLY, 0o644)
	if err != nil {
		log.Error().Err(err).Msg("failed to open stdiotest.yaml")
		os.Exit(1)
	}

	var spec Spec
	if err := yaml.NewDecoder(f).Decode(&spec); err != nil {
		log.Error().Err(err).Msg("failed to decode stdiotest.yaml")
		os.Exit(1)
	}

	return spec
}
