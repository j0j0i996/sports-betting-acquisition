package main

import (
	"db/client"
	model "db/model"
	"fdu/acquisition"
	"fdu/transform"

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

	// acquire, transform and load data
	acquired_fixtures := acquisition.GetHistoricData("2223", "Bundesliga")
	for _, fixture := range acquired_fixtures {
		historic_odds := transform.HistoricOddsSourceModelToDbModel(fixture)
		/*
			db_client.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "fixture_id"}},
				UpdateAll: true,
			}).Create(&historic_odds)
		*/
		// TODO create composite index on bookmaker
		// TODO round before writing
		db_client.Create(&historic_odds)
	}
}
