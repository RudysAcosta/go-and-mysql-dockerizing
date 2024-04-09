package config

import "os"

// Config structure to store the application configuration
type Config struct {
	Port     string
	Database DatabaseConfig
}

// DatabaseConfig structure to store database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

// NewConfig creates a new instance of Config with default values
func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("PORT"),
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
