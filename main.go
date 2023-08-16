package main

import (
	"unjuk_keterampilan/app/configs"
	"unjuk_keterampilan/app/databases"
	"unjuk_keterampilan/app/logger"
	"unjuk_keterampilan/app/wire"

	"github.com/spf13/viper"
)

func main() {
	logger.SetupLogger()
	configs.SetupViper()
	databases.NewConnection()

	e := wire.SetupApp()
	e.Logger.Fatal(e.Start(":" + viper.GetString("app.port")))
}
