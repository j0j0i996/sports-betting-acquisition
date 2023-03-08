package main

import (
	"fmt"
	"log"
	"frf/acquisition/leagues"
	"frf/acquisition/fixtures"
)

func main() {
	id, err := leagues.GetLeagueId("Bundesliga", "Austria", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	data := fixtures.GetFixtures(id, 2022, false)
	fmt.Println(data[0].Teams.Home.Name)
}