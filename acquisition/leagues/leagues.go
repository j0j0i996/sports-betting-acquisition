package leagues

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
	"errors"
	"path/filepath"
	"frf/acquisition/football_api"
)

type Response struct {
	LeagueList []LeagueItem `json:"response"`
}

type LeagueItem struct {
	League League
	Country Country
}

type League struct {
    Id int
    Name string
}

type Country struct {
    Name string
    Code string
}

// GetLeages TODO
func getLeagues(from_local bool) []LeagueItem {
	p := filepath.Join("local_datasets", "raw", "leagues.json")
	if !from_local {
		// Get Data from API
		raw_data := football_api.GetData("leagues")
		ioutil.WriteFile(p, raw_data, 0644)
	}
	var res Response
	// TODO local folder
	var jsonFile, err = os.Open(p)
		if err != nil {
			log.Fatal(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &res)
	return res.LeagueList
}

func GetLeagueId(name string, country string, from_local bool) (int, error) {
	league_list := getLeagues(from_local)
	for _, item := range league_list {
		if item.League.Name == name && item.Country.Name == country {
			return item.League.Id, nil
		}
	}
	return 0, errors.New("No id found for name and country")
}