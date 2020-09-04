// Package models gorm's models
package models

import (
	"github.com/103cuong/gorm_kit/configs"
	// postgres driver
	_ "gorm.io/driver/postgres"
)

// GetCats fetch all cats.
func GetCats(cats *[]Cat) (err error) {
	err = configs.DB.Find(&cats).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateCat create new cat.
func CreateCat(cat *Cat) (err error) {
	err = configs.DB.Create(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCatByID fetch cat by ID.
func GetCatByID(cat *Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCat update cat.
func UpdateCat(cat *Cat) (err error) {
	err = configs.DB.Save(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCat delete cat.
func DeleteCat(cat *Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
