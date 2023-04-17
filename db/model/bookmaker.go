package model

type Bookmaker struct {
	Id   uint   `gorm:"primary_key"`
	Name string `gorm:"not null; uniqueIndex"`
	Slug string `gorm:"not null"`
}
