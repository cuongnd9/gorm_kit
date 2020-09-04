package main

import (
	"github.com/103cuong/gorm_kit/configs"
	"github.com/103cuong/gorm_kit/models"
	"github.com/103cuong/gorm_kit/routes"
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

	// migrate the scheme.
	err = configs.DB.AutoMigrate(&models.Cat{})
	if err != nil {
		panic("failed to migrate the schema")
	}

	router := routes.SetupRouter()
	err = router.Run()
	if err != nil {
		panic("failed to run server")
	}
}
