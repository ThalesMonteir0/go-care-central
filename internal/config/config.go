package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv           string `mapstructure:"APP_ENV"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`
}

func LoadVariables() *Env {
	var env Env

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("não foi possivel ler as variaveis de ambiente. Err: ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("não foi possivel ler as variaveis de ambiente. Err: ", err)
	}

	if env.AppEnv == "DEV" {
		println("App rodando no ambiente de desenvolvimento.")
	}

	return &env
}
