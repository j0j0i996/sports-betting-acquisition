package acquisition

import (
	"encoding/json"
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
