package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func SetupViper() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	log.Info().Msg("Success load configuration application...")
}
