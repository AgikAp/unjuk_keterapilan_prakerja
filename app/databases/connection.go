package databases

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.pass"),
		viper.GetString("database.dbname"),
		viper.GetString("database.port"),
		viper.GetString("database.sslmode"),
		viper.GetString("database.timezone"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Failed connection database : " + err.Error())
	}

	log.Info().Msg("Success connection database")
	return db
}
