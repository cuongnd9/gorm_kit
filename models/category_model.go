// Package models gorm's models
package models

// Category category model struct
type Category struct {
	Pure `gorm:"embedded"`
	Name string `json:"name"`
}

// TableName get table name
func (category *Category) TableName() string {
	return "categories"
}
