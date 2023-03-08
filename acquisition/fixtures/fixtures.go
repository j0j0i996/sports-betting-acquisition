package fixtures

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"time"
	"sort"
	"frf/acquisition/football_api"
)

type Response struct {
	FixtureList []FixtureItem `json:"response"`
}

type FixtureItem struct {
	Meta Meta `json:"fixture"`
	Teams FixtureTeams
	Goals Goals
}

type Meta struct {
    Id int
    Date time.Time
	Status Status
}

type FixtureTeams struct {
    Home Team
    Away Team
}

type Team struct {
    Id int
    Name string
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
func GetFixtures(league_id int, season int, from_local bool) (map[int]FixtureItem, []FixtureItem) {
	p := filepath.Join("local_datasets", "raw", "fixtures_" +  fmt.Sprint(league_id) + "_" + fmt.Sprint(season) + ".json")
	if !from_local {
		// Get Data from API
		var parameter_map = map[string]string{"season": fmt.Sprint(season), "league": fmt.Sprint(league_id)}
		raw_data := football_api.GetData("fixtures", parameter_map)
		ioutil.WriteFile(p, raw_data, 0644)
	}
	var res Response
	var jsonFile, err = os.Open(p)
		if err != nil {
			log.Fatal(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &res)

	fixture_list := res.FixtureList

	// sort by date
	sort.SliceStable(fixture_list, func(i, j int) bool {
		return fixture_list[i].Meta.Date.Before(fixture_list[j].Meta.Date)
	})

	fixture_map := make(map[int]FixtureItem)

	for _, item := range fixture_list {
		fixture_map[item.Meta.Id] = item
	}	
	return fixture_map, fixture_list
}