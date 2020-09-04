package models

import (
	"github.com/103cuong/gorm_kit/configs"
	_ "gorm.io/driver/postgres"
)

// Fetch all cats.
func GetCats(cats *[]Cat) (err error) {
	err = configs.DB.Find(&cats).Error
	if err != nil {
		return err
	}
	return nil
}

// Create new cat.
func CreateCat(cat *Cat) (err error) {
	err = configs.DB.Create(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// Fetch cat by ID.
func GetCatByID(cat *Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// Update cat.
func UpdateCat(cat *Cat) (err error) {
	err = configs.DB.Save(&cat).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete cat.
func DeleteCat(cat *Cat, id string) (err error) {
	err = configs.DB.Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
