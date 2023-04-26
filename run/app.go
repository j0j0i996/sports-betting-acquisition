package main

import (
	"db/insert"
	"db/model"
	fas_sync "fas/sync"
	fdu_sync "fdu/sync"
)

func main() {

	boolInsertBookmaker := false
	boolSyncLeagues := false
	boolSyncTeams := false
	boolSyncSportsApi := true
	boolSyncDataUk := false

	// Insert bookmakers
	if boolInsertBookmaker {
		insert.InsertBookmaker("bet365", "B365")
	}

	seasons := [3]string{
		"2020/21",
		"2021/22",
		"2022/23",
	}
	leagues := [3]model.League{
		{Name: "Bundesliga", Country: "Germany"},
		{Name: "2. Bundesliga", Country: "Germany"},
		{Name: "Premier League", Country: "England"},
		//{Name: "Serie A", Country: "Italy"},
		//{Name: "Eredivisie", Country: "Netherlands"},
	}
	// problem with eredivisie fixture time heerenveen vs willem ii and serie a paris vs saint germain

	// prozess Leagues
	if boolSyncLeagues {
		fas_sync.SyncLeagues()
	}

	for _, league := range leagues {
		for _, season := range seasons {
			// synchronize football api sports
			if boolSyncSportsApi {
				fas_sync.Sync(league, season, boolSyncTeams)
			}

			// synchronize football data uk
			if boolSyncDataUk {
				fdu_sync.Sync(league, season)
			}
		}
	}

}
