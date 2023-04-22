package transform

import (
	db_getters "db/getters"
	"db/model"
	db_model "db/model"
	acquisition "fdu/acquisition"
	"reflect"

	"github.com/pkg/errors"
)

func HistoricOddsSourceModelToDbModel(source_fixture_item acquisition.FixtureItem) ([]db_model.HistoricOdds, error) {

	// get team ids:
	home, err := db_getters.GetTeamFromSimilarTeamName(source_fixture_item.HomeTeam)
	if err != nil {
		return []db_model.HistoricOdds{}, errors.Wrap(err, "finding similar team name for: "+source_fixture_item.HomeTeam)
	}
	away, err := db_getters.GetTeamFromSimilarTeamName(source_fixture_item.AwayTeam)
	if err != nil {
		return []db_model.HistoricOdds{}, errors.Wrap(err, "finding similar team name for: "+source_fixture_item.AwayTeam)
	}

	// find fixture
	fixture, err := db_getters.FindFixtureFromTeamIdsAndTimestampInDB(source_fixture_item.Date, home, away)
	if err != nil {
		return []db_model.HistoricOdds{}, errors.Wrap(err, "did not close fixture time similar for "+home.Name+" vs. "+away.Name)
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

	return historic_odds, nil
}
