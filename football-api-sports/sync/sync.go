package sync

import (
	"db/client"
	"db/getters"
	"db/model"
	"fas/acquisition"
	"fas/transform"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

var db_client = client.GetClient()

func formatSeason(season string) uint {
	yearString := strings.Split(season, "/")[0]
	yearInt, err := strconv.Atoi(yearString)
	if err != nil {
		panic(err)
	}
	return uint(yearInt)
}

func SyncLeagues() {
	// fetches all leagues from a specific year and stores them in db

	// Get Leagues as list
	league_list := acquisition.GetLeagues()

	// transform and insert all fixtuers
	for _, league := range league_list {
		insertLeague := transform.LeagueApiModelToDbModel(league)

		// Insert (Do nothing on conflict)
		db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&insertLeague)
	}
}

func syncTeams(season uint, league_id uint) {
	// fetches all leagues from a specific year and stores them in db

	// Get Leagues as list
	team_list := acquisition.GetTeams(league_id, season)

	// transform and insert all fixtuers
	for _, team := range team_list {
		insertTeam := transform.TeamApiModelToDbModel(team)

		// Insert (Do nothing on conflict)
		db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&insertTeam)
	}
}

func syncFixtures(season uint, league_id uint) {
	// fetches all fixtures from a specific year and league and stores them in db

	// Get Fixtures as list
	fixture_list := acquisition.GetFixtures(league_id, season)

	// transform and insert all fixtuers
	for _, fixture := range fixture_list {
		insertFixture := transform.FixtureApiModelToDbModel(fixture, season)

		// Insert (Update time, goals and result on conflict)
		db_client.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"time", "home_team_goals", "away_team_goals", "result"}),
		}).Create(&insertFixture)
	}
}

func Sync(input_league model.League, season string, boolSyncTeams bool) {

	league := getters.GetLeagueByNameAndCountry(input_league.Name, input_league.Country)

	year := formatSeason(season)
	if boolSyncTeams {
		syncTeams(year, league.Id)
	}

	//sync fixtures
	syncFixtures(year, league.Id)

	// avoid too many requests error
	time.Sleep(3 * time.Second)
}
