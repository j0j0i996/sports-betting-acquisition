package transform

import (
	model "db/model"
	acquisition "fas/acquisition"
)

func TeamApiModelToDbModel(api_team acquisition.TeamItem) model.Team {

	// Build db model
	return model.Team{
		Id:      api_team.Team.Id,
		Name:    api_team.Team.Name,
		Country: api_team.Team.Country,
	}

}
