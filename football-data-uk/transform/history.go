package transform

import (
	db_getters "db/getters"
	"db/model"
	db_model "db/model"
	acquisition "fdu/acquisition"
	"log"
	"reflect"
)

func HistoricOddsSourceModelToDbModel(source_fixture_item acquisition.FixtureItem) []db_model.HistoricOdds {

	// get team ids:
	home, err := db_getters.GetTeamFromSimilarTeamName(source_fixture_item.HomeTeam)
	if err != nil {
		log.Fatal("did not find similar team name: " + source_fixture_item.HomeTeam + ". Error: " + err.Error())
	}
	away, err := db_getters.GetTeamFromSimilarTeamName(source_fixture_item.AwayTeam)
	if err != nil {
		log.Fatal("did not find similar team name " + source_fixture_item.AwayTeam + ". Error: " + err.Error())
	}

	// find fixture
	fixture, err := db_getters.FindFixtureFromTeamIdsAndTimestampInDB(source_fixture_item.Date, home, away)
	if err != nil {
		log.Fatal("did not close fixture time similar for " + home.Name + " vs. " + away.Name + ". Error: " + err.Error())
	}

	// convert to one entry per bookmaker
	var historic_odds []model.HistoricOdds
	for _, bookmaker_slug := range acquisition.INTEGRATED_BOOKMAKERS {
		// find bookmaker:
		bookmaker := db_getters.GetBookmakersBySlugInDB(bookmaker_slug)
		new_entry := model.HistoricOdds{
			FixtureId:   fixture.Id,
			BookmakerId: bookmaker.Id,
			HomeOdds:    reflect.ValueOf(source_fixture_item).FieldByName(bookmaker_slug + "H").Float(),
			AwayOdds:    reflect.ValueOf(source_fixture_item).FieldByName(bookmaker_slug + "A").Float(),
			DrawOdds:    reflect.ValueOf(source_fixture_item).FieldByName(bookmaker_slug + "D").Float(),
		}
		historic_odds = append(historic_odds, new_entry)
	}

	return historic_odds
}
