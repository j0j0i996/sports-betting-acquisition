package transform

import (
	acquisition "fas/acquisition"
	db "fas/db"
)

func TeamApiModelToDbModel(api_team acquisition.TeamItem) db.Team {

	// Build db model
	return db.Team{
		Id:      api_team.Team.Id,
		Name:    api_team.Team.Name,
		Country: api_team.Team.Country,
	}

}
