package db

type League struct {
	Id      uint `gorm:"primary_key"`
	Name    string
	Country string
}
