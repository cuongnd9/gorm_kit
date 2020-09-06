package main

import (
	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	configs.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  configs.BuildDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate database.
	migrations.MigrateDB("up")

	// your code: gin, mux, grpc, etc.
}
