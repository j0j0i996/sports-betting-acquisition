package transform

import (
	acquisition "fas/acquisition"
	db "fas/db"
)

func FixtureApiModelToDbModel(api_fixture acquisition.FixtureItem) db.Fixture {

	// determine result
	var result db.Result
	if api_fixture.Goals.Home.Get() == nil {
		result = db.TBD
	} else if *api_fixture.Goals.Home.Get() > *api_fixture.Goals.Away.Get() {
		result = db.Home
	} else if *api_fixture.Goals.Home.Get() == *api_fixture.Goals.Away.Get() {
		result = db.Draw
	} else {
		result = db.Away
	}

	// Build db model
	return db.Fixture{
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
