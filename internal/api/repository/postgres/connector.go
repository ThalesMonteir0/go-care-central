package postgres

import (
	"database/sql"
	"fmt"
	"github.com/ThalesMonteir0/go-care-central/internal/config"
	"log"
)

func NewPostgresConnection(env config.Env) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.PostgresHost,
		env.PostgresPort,
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresDB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("nao foi possivel iniciar o banco de dados - postgres")
	}

	return db
}
