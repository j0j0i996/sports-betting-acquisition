package main

import (
	"fmt"
	"time"

	acquisition "fas/acquisition"
	transform "fas/transform"

	"db/client"
	db_getters "db/getters"

	"gorm.io/gorm/clause"
)

// Connect to DB
var db_client = client.GetClient()

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

	boolSyncLeagues := false
	boolSyncTeams := true

	// prozess Leagues
	if boolSyncLeagues {
		syncLeagues()
	}

	// prozess Fixturesteam
	years := [2]int{2021, 2022}

	leagues := [5]LeagueDescription{
		{Name: "Bundesliga", Country: "Germany"},
		{Name: "2. Bundesliga", Country: "Germany"},
		{Name: "Premier League", Country: "England"},
		{Name: "Serie A", Country: "Italy"},
		{Name: "Eredivisie", Country: "Netherlands"},
	}

	for _, league := range leagues {
		for _, year := range years {
			// Get League
			league := db_getters.GetLeagueByNameAndCountry(league.Name, league.Country)
			fmt.Println(league)

			// sync teams if desired
			if boolSyncTeams {
				syncTeams(year, league.Id)
			}

			//sync fixtures
			syncFixtures(year, league.Id)

			// avoid too many error
			time.Sleep(10 * time.Second)
		}
	}

}
