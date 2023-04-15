package model

import "fmt"

type League struct {
	Id      uint `gorm:"primary_key"`
	Name    string
	Country string
}

func getLeagueName(league League) string {
	fmt.Println(league.Name)
	return league.Name
}
