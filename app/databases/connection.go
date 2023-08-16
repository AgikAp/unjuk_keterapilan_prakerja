package databases

import (
	"fmt"
	"unjuk_keterampilan/app/modules/notes/model"
	userModel "unjuk_keterampilan/app/modules/users/model"

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

	log.Info().Msg("Try to migration...")
	if err = Migration(db); err != nil {
		log.Fatal().Msg("Failed migration : " + err.Error())
	}

	Seeder(db)
	log.Info().Msg("Success connection database")
	return db
}

func Migration(db *gorm.DB) (err error) {
	return db.AutoMigrate(model.Note{}, userModel.User{})
}

func Seeder(db *gorm.DB) {
	var count int64
	db.Model(userModel.User{}).Where("username = ?", "system").Count(&count)
	if count == 0 {
		user := userModel.User{}
		user.Username = "system"
		user.Email = "system@mail.com"
		db.Model(userModel.User{}).Create(&user)
	}
}
