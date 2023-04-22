package main

import (
	"db/client"
	model "db/model"
	"fdu/acquisition"
	"fdu/transform"
	"fmt"
	"log"

	"gorm.io/gorm/clause"
)

var db_client = client.GetClient()

func main() {

	// Insert bookmakers
	bookmakers := [1]model.Bookmaker{
		{Name: "bet365", Slug: "B365"},
	}
	db_client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&bookmakers)

	years := [3]string{
		//"1617",
		//"1718",
		//"1819",
		//"1920",
		"2021",
		"2122",
		"2223",
	}
	leagues := [5]string{"Bundesliga", "2. Bundesliga", "Premier League", "Eredivisie", "Serie A"}

	for _, league := range leagues {
		for _, year := range years {

			acquired_fixtures, err := acquisition.GetHistoricData(year, league)
			if err != nil {
				log.Fatal("error in data acquisition" + err.Error())
			}
			fmt.Println("Received Data")
			fmt.Println(acquired_fixtures)
			for _, fixture := range acquired_fixtures {
				historic_odds, err := transform.HistoricOddsSourceModelToDbModel(fixture)
				if err != nil {
					log.Fatal("error in data acquisition" + err.Error())
				}
				fmt.Println("Inserting to Database: ")
				fmt.Println(historic_odds)
				db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&historic_odds)

			}
		}
	}

}
