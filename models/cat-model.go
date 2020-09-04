package models

type Cat struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (cat *Cat) TableName() string {
	return "cat"
}
