package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rudysacosta/go-and-mysql-dockerizing/config"
)

// Database struct to hold the connection
type Database struct {
	*sqlx.DB
}

// NewDatabase creates a new instance of the Database struct
func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.DB.Close()
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (db *Database) Ping() error {
	return db.DB.Ping()
}
