// Package configs database configs
package configs

import (
	"fmt"

	"gorm.io/gorm"
)

// DB GORM DB instance
var DB *gorm.DB

// DBConfig database configuration
type DBConfig struct {
	Host     string
	Port     uint
	User     string
	DBName   string
	Password string
}

// BuildDBConfig build DB config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     5432,
		User:     "postgres",
		DBName:   "postgres",
		Password: "postgres",
	}
	return &dbConfig
}

// BuildDSN build Postgresql's DSN
func BuildDSN() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
