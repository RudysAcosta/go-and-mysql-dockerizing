package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rudysacosta/go-and-mysql-dockerizing/config"
)

// Database struct to hold the connection
type Database struct {
	*sql.DB
}

// NewDatabase creates a new instance of the Database struct
func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	return db.DB.Close()
}
