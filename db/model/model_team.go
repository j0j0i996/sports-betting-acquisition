package model

import "fmt"

type Team struct {
	Id      uint `gorm:"primary_key"`
	Name    string
	Country string
}

func getTeamName(team Team) string {
	fmt.Println(team.Name)
	return team.Name
}
