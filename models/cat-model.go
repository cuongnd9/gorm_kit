package models

// Cat cat model struct
type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// TableName get table name
func (cat *Cat) TableName() string {
	return "cat"
}
