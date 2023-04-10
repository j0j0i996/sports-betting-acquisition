package main

import (
	acquisition "fas/acquisition"
	db "fas/db"
	transform "fas/transform"
	"log"

	"gorm.io/gorm/clause"
)

// Connect to DB
var db_client = db.GetClient()

type LeagueDescription struct {
	Name    string
	Country string
}

func syncFixtures(year int, league LeagueDescription) {
	// fetches all fixtures from a specific year and league and stores them in db

	// Get League ID
	id, err := acquisition.GetLeagueId(league.Name, league.Country)
	if err != nil {
		log.Fatal(err)
	}

	// Get Fixtures as list
	fixture_list := acquisition.GetFixtures(id, year)

	// transform and insert all fixtuers
	for _, fixture := range fixture_list {
		insertFixture := transform.FixtureApiModelToDbModel(fixture)

		// Insert (Update time, goals and result on conflict)
		db_client.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"time", "home_team_goals", "away_team_goals", "result"}),
		}).Create(&insertFixture)
	}

}

func syncLeagues() {
	// fetches all leagues from a specific year and stores them in db

	// Get Leagues as list
	league_list := acquisition.GetLeagues()

	// transform and insert all fixtuers
	for _, league := range league_list {
		insertLeague := transform.LeagueApiModelToDbModel(league)

		// Insert (Do nothing on conflict)
		db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&insertLeague)
	}

}

func main() {

	// prozess Leagues
	syncLeagues()

	// prozess Fixtures
	years := [1]int{2022}
	leagues := [1]LeagueDescription{{Name: "Serie A", Country: "Italy"}}

	for _, year := range years {
		for _, league := range leagues {
			syncFixtures(year, league)
		}
	}

}
