package configs

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Database configuration.
type DBConfig struct {
	Host     string
	Port     uint
	User     string
	DBName   string
	Password string
}

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

func BuildDSN() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
