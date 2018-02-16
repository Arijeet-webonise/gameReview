package database

import (
	"database/sql"

	"github.com/knq/dburl"
)

type IDatabaseConnection interface {
	Initialise(map[string]string) (*sql.DB, error)
}
type DatabaseWrapper struct {
	DB *sql.DB
}

func (dw *DatabaseWrapper) Initialise(config map[string]string) (*sql.DB, error) {
	// need to construct this URL from the params map passed
	connStr := config["driver"] + "://" + config["user"] + ":" + config["password"] + "@" + config["host"] + "/" + config["db"] + "?sslmode=disable"
	u, err := dburl.Open(connStr)

	return u, err
}
