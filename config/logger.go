package config

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitLogger() {
	if viper.GetString("ENV") == "test" {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var writer io.Writer
	var subWriters []io.Writer = []io.Writer{}

	if len(subWriters) >= 1 {
		writer = zerolog.MultiLevelWriter(subWriters...)
	} else {
		writer = os.Stdout
	}

	log.Logger = log.Output(writer)
	log.Info().Msg("Logger initialized")
}
