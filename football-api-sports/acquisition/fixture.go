package acquisition

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tee8z/nullable"
)

type FixtureResponse struct {
	FixtureList []FixtureItem `json:"response"`
}

type FixtureItem struct {
	Meta   Meta `json:"fixture"`
	Teams  FixtureTeams
	Goals  Goals
	League FixtureLeague
}

type Meta struct {
	Id     uint
	Date   time.Time
	Status Status
}

type FixtureLeague struct {
	Id uint
}

type FixtureTeams struct {
	Home Team
	Away Team
}

type Team struct {
	Id     uint
	Name   string
	Winner bool
}

type Goals struct {
	Home nullable.Int8
	Away nullable.Int8
}

type Status struct {
	Long string
}

// GetLeages TODO
func GetFixtures(league_id uint, season int) []FixtureItem {
	var res FixtureResponse
	// Get Data from API
	var parameter_map = map[string]string{"season": fmt.Sprint(season), "league": fmt.Sprint(league_id)}
	raw_data := GetData("fixtures", parameter_map)
	json.Unmarshal(raw_data, &res)

	fixture_list := res.FixtureList

	return fixture_list
}
