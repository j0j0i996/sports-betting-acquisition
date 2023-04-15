package acquisition

import (
	"encoding/json"
	"errors"
)

type LeagueResponse struct {
	LeagueList []LeagueItem `json:"response"`
}

type LeagueItem struct {
	League  League
	Country Country
}

type League struct {
	Id   uint
	Name string
}

type Country struct {
	Name string
	Code string
}

// GetLeages TODO
func GetLeagues() []LeagueItem {
	// Get Data from API
	raw_data := GetData("leagues")
	var res LeagueResponse
	json.Unmarshal(raw_data, &res)

	return res.LeagueList
}

func GetLeagueId(name string, country string) (uint, error) {
	league_list := GetLeagues()
	for _, item := range league_list {
		if item.League.Name == name && item.Country.Name == country {
			return item.League.Id, nil
		}
	}
	return 0, errors.New("no id found for name and country")
}
