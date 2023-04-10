package transform

import (
	acquisition "fas/acquisition"
	db "fas/db"
)

func LeagueApiModelToDbModel(api_league acquisition.LeagueItem) db.League {

	// Build db model
	return db.League{
		Id:      api_league.League.Id,
		Name:    api_league.League.Name,
		Country: api_league.Country.Name,
	}

}
