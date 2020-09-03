package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Product model.
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=0.0.0.0 user=postgres password=postgres dbname=postgres port=5432",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate the schema.
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to migrate the schema")
	}

	// create.
	db.Create(&Product{
		Code:  "ABC_XYZ",
		Price: 100,
	})

	// read.
	var product Product
	// find product with integer primary key.
	db.First(&product, 1)
	fmt.Printf("read 1: %v\n", product)
	// find product with code ABC_XYZ
	db.First(&product, "code = ?", "ABC_XYZ")
	fmt.Printf("read 2: %v\n", product)

	// update.
	// update product's price to 120.
	db.Model(&product).Update("Price", 120)
	// update multiple fields.
	db.Model(&product).Updates(Product{
		Code:  "MOEW_MOEW",
		Price: 150,
	})
	// map[string]interface{} in Go: https://bitfieldconsulting.com/golang/map-string-interface.
	db.Model(&product).Updates(map[string]interface{}{"Price": 399, "Code": "KAKAKA"})

	// delete.
	db.Delete(&product, 1)
}
