package transform

import (
	model "db/model"
	acquisition "fas/acquisition"
)

func LeagueApiModelToDbModel(api_league acquisition.LeagueItem) model.League {

	// Build db model
	return model.League{
		Id:      api_league.League.Id,
		Name:    api_league.League.Name,
		Country: api_league.Country.Name,
	}

}
