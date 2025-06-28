package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppEnv           string `json:"APP_ENV"`
	PostgresHost     int    `json:"POSTGRES_HOST"`
	PostgresUser     string `json:"POSTGRES_USER"`
	PostgresPassword string `json:"POSTGRES_PASSWORD"`
	PostgresDB       string `json:"POSTGRES_DB"`
	PostgresPort     string `json:"POSTGRES_PORT"`
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
