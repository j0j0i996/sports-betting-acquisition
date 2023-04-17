package transform

import (
	db_getters "db/getters"
	"db/model"
	db_model "db/model"
	acquisition "fdu/acquisition"
	"fmt"
	"log"
)

func HistoricOddsSourceModelToDbModel(historic_odds acquisition.FixtureItem) db_model.HistoricOdds {

	// get team ids:
	home, err := db_getters.GetTeamFromSimilarTeamName(historic_odds.HomeTeam)
	if err != nil {
		log.Fatal("did not find similar team name")
	}
	away, err := db_getters.GetTeamFromSimilarTeamName(historic_odds.AwayTeam)
	if err != nil {
		log.Fatal("did not find similar team name")
	}

	// find fixture
	fixture, err := db_getters.FindFixtureFromTeamIdsAndTimestampInDB(historic_odds.Time, home, away)
	if err != nil {
		log.Fatal("did not find similar team name")
	}

	fmt.Println(fixture)

	// Build db model
	/*
		return model.Fixture{
			Id:            api_fixture.Meta.Id,
			Time:          api_fixture.Meta.Date,
			HomeTeamId:    api_fixture.Teams.Home.Id,
			AwayTeamId:    api_fixture.Teams.Away.Id,
			HomeTeamGoals: api_fixture.Goals.Home,
			AwayTeamGoals: api_fixture.Goals.Away,
			Result:        result,
			LeagueId:      api_fixture.League.Id,
		}
	*/
	return model.HistoricOdds{}
}
