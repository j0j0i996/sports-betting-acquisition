package db

type Team struct {
	Id      uint `gorm:"primary_key"`
	Name    string
	Country string
}
