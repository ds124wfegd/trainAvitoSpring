package storagePostgres

import (
	"fmt"

	"github.com/ds124wfegd/trainAvitoSpring/config"
	"github.com/jmoiron/sqlx"
)

func InitPsqlDB(conf *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s, user=%s, password=%s, dbname=%s, sslmode=%s",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.DBName, conf.Postgres.SSLMode)
	return sqlx.Connect(conf.Postgres.PgDriver, connectionUrl)
}
