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

func syncFixtures(year int, league_id uint) {
	// fetches all fixtures from a specific year and league and stores them in db

	// Get Fixtures as list
	fixture_list := acquisition.GetFixtures(league_id, year)

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

func syncTeams(year int, league_id uint) {
	// fetches all leagues from a specific year and stores them in db

	// Get Leagues as list
	team_list := acquisition.GetTeams(league_id, year)

	// transform and insert all fixtuers
	for _, team := range team_list {
		insertTeam := transform.TeamApiModelToDbModel(team)

		// Insert (Do nothing on conflict)
		db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&insertTeam)
	}

}

func main() {

	boolSyncLeagues := true
	boolSyncTeams := true

	// prozess Leagues
	if boolSyncLeagues {
		syncLeagues()
	}

	// prozess Fixturesteam
	years := [1]int{2022}
	leagues := [1]LeagueDescription{{Name: "Serie A", Country: "Italy"}}

	for _, year := range years {
		for _, league := range leagues {
			// Get League ID
			league_id, err := acquisition.GetLeagueId(league.Name, league.Country)
			if err != nil {
				log.Fatal(err)
			}
			syncTeams(year, league_id)
			if boolSyncTeams {
				syncTeams(year, league_id)
			}
			syncFixtures(year, league_id)
		}
	}

}
