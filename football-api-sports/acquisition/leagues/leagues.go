package leagues

import (
	"encoding/json"
	"errors"
	"fas/acquisition/football_api"
)

type Response struct {
	LeagueList []LeagueItem `json:"response"`
}

type LeagueItem struct {
	League  League
	Country Country
}

type League struct {
	Id   int
	Name string
}

type Country struct {
	Name string
	Code string
}

// GetLeages TODO
func getLeagues() []LeagueItem {
	// Get Data from API
	raw_data := football_api.GetData("leagues")
	var res Response
	json.Unmarshal(raw_data, &res)

	return res.LeagueList
}

func GetLeagueId(name string, country string) (int, error) {
	league_list := getLeagues()
	for _, item := range league_list {
		if item.League.Name == name && item.Country.Name == country {
			return item.League.Id, nil
		}
	}
	return 0, errors.New("no id found for name and country")
}
