package fixtures

import (
	"encoding/json"
	"fas/acquisition/football_api"
	"fmt"
	"sort"
	"time"
)

type Response struct {
	FixtureList []FixtureItem `json:"response"`
}

type FixtureItem struct {
	Meta  Meta `json:"fixture"`
	Teams FixtureTeams
	Goals Goals
}

type Meta struct {
	Id     uint
	Date   time.Time
	Status Status
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
	Home int
	Away int
}

type Status struct {
	Long string
}

// GetLeages TODO
func GetFixtures(league_id int, season int) []FixtureItem {
	var res Response
	// Get Data from API
	var parameter_map = map[string]string{"season": fmt.Sprint(season), "league": fmt.Sprint(league_id)}
	raw_data := football_api.GetData("fixtures", parameter_map)
	json.Unmarshal(raw_data, &res)

	fixture_list := res.FixtureList

	// sort by date
	sort.SliceStable(fixture_list, func(i, j int) bool {
		return fixture_list[i].Meta.Date.Before(fixture_list[j].Meta.Date)
	})

	return fixture_list
}
