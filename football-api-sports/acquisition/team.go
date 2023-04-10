package acquisition

import (
	"encoding/json"
	"fmt"
)

type TeamsResponse struct {
	TeamList []TeamItem `json:"response"`
}

type TeamItem struct {
	Team Team
}

type Team struct {
	Id      uint
	Name    string
	Country string
}

// GetLeages TODO
func GetTeams(league_id uint, season int) []TeamItem {
	var res TeamsResponse
	// Get Data from API
	var parameter_map = map[string]string{"season": fmt.Sprint(season), "league": fmt.Sprint(league_id)}
	raw_data := GetData("teams", parameter_map)
	json.Unmarshal(raw_data, &res)
	fmt.Println(res)

	return res.TeamList
}
