package main

import (
	"frf/acquisition/leagues"
	"frf/acquisition/fixtures"
	"frf/preperation"
	"fmt"
	"log"
)

func main() {
	// Get League ID
	id, err := leagues.GetLeagueId("Bundesliga", "Germany", true)
	if err != nil {
		log.Fatal(err)
	}

	// Get Fixtures as map (indexed by fixture id) and as list
	fixture_map, fixture_list := fixtures.GetFixtures(id, 2022, true)

	// Calculate Stats
	pre_game_stats_map := preperation.GetPregameStatsMap(fixture_list, fixture_map)

	fmt.Println(pre_game_stats_map)
}