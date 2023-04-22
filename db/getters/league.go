package getters

import model "db/model"

func GetLeagueByNameAndCountry(name string, country string) model.League {
	var league model.League
	db_client.Where("name = ? AND country = ?", name, country).First(&league)
	return league
}
