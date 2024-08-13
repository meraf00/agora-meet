package config

import (
	"log"

	"github.com/spf13/viper"
)

type mongodbConfig struct {
	DbName string
	Uri    string
}

var AppConfig = struct {
	Port    int
	Mongodb mongodbConfig
}{}

func LoadConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading config file %s", err)
	}

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Mongodb.Uri = viper.GetString("MONGODB_URI")
	AppConfig.Mongodb.DbName = viper.GetString("MONGODB_DATABASE_NAME")
}
