package transform

import (
	acquisition "fas/acquisition"

	model "db/model"
)

func FixtureApiModelToDbModel(api_fixture acquisition.FixtureItem) model.Fixture {

	// determine result
	var result model.Result
	if api_fixture.Goals.Home.Get() == nil {
		result = model.TBD
	} else if *api_fixture.Goals.Home.Get() > *api_fixture.Goals.Away.Get() {
		result = model.Home
	} else if *api_fixture.Goals.Home.Get() == *api_fixture.Goals.Away.Get() {
		result = model.Draw
	} else {
		result = model.Away
	}

	// Build db model
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

}
