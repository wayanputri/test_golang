package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DBURL        string
	DBName       string
	DBCollection string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("URL"); found {
		app.DBURL = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBCOLL"); found {
		app.DBCollection = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config: ", err.Error())
			return nil
		}

		app.DBURL = viper.Get("URL").(string)
		app.DBName = viper.Get("DBNAME").(string)
		app.DBCollection = viper.Get("DBCOLL").(string)
	}
	return &app
}
