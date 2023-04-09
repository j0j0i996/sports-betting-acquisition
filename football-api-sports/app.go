package main

import (
	"fas/acquisition/fixtures"
	"fas/acquisition/leagues"
	"fas/interfaces"
	"log"
)

// Connect to DB
var db = interfaces.GetClient()

type LeagueDescription struct {
	Name    string
	Country string
}

func sync_fixtures(year int, league LeagueDescription) {
	// fetches all fixtures from a specific year and league and stores it in db

	// Get League ID
	id, err := leagues.GetLeagueId(league.Name, league.Country)
	if err != nil {
		log.Fatal(err)
	}

	// Get Fixtures as list
	fixture_list := fixtures.GetFixtures(id, year)

	// iterate over all fixtures
	for _, fixture := range fixture_list {

		// determine result
		var result interfaces.Result
		if fixture.Goals.Home > fixture.Goals.Away {
			result = interfaces.Home
		} else if fixture.Goals.Home == fixture.Goals.Away {
			result = interfaces.Draw
		} else {
			result = interfaces.Away
		}

		// insert to db
		insertFixture := interfaces.Fixture{
			Id:         fixture.Meta.Id,
			Time:       fixture.Meta.Date,
			HomeTeamId: fixture.Teams.Home.Id,
			AwayTeamId: fixture.Teams.Away.Id,
			Result:     result,
		}

		db.Create(&insertFixture)
	}

}

func main() {

	years := [5]int{2018, 2019, 2020, 2021, 2022}
	leagues := [1]LeagueDescription{{Name: "Bundesliga", Country: "Germany"}}

	for _, year := range years {
		for _, league := range leagues {
			sync_fixtures(year, league)
		}
	}

}
