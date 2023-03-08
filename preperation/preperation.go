package preperation

import (
	"frf/acquisition/leagues"
	"frf/acquisition/fixtures"
	"encoding/csv"
	"log"
	"path/filepath"
	"os"
	"fmt"
)

type FixturePreGameStats struct {
	Id int
	Home TeamPreGameStats
	Away TeamPreGameStats
}

type TeamPreGameStats struct {
	LastXFixtureIds []int
	GamesAnalzed int
	Wins int
	Draws int
	Losses int
}

func getIdsOfLastXFixturesOfTeam(fixture_position int, team_id int, fixture_list []fixtures.FixtureItem, fixtures_to_be_found int) []int {
	var ids_of_last_fixtures_of_team []int
	fixtures_found := 0
	// Iterate through reversed fixture list until X fixtures are found
	for i := fixture_position-1; i >= 0 && fixtures_found < fixtures_to_be_found; i-- {

		// Continue if game not played
		if fixture_list[i].Meta.Status.Long == "Not Started" || fixture_list[i].Meta.Status.Long == "Time to be defined"{
			continue
		}

		// append fixture if Home or Away Team equals Team
		if fixture_list[i].Teams.Home.Id == team_id || fixture_list[i].Teams.Away.Id == team_id {
			ids_of_last_fixtures_of_team = append(ids_of_last_fixtures_of_team, fixture_list[i].Meta.Id)
			fixtures_found++
		}
	} 
	return ids_of_last_fixtures_of_team
}

// Calculate how many fixtures in a given set were won, lost, drew by a given team
func getWdlStats(ids_of_last_fixtures_of_team []int, team_id int, fixture_map map[int]fixtures.FixtureItem) (int, int, int, int) {
	analyzed := 0; wins := 0; draws := 0; losses := 0;
	for _, id := range ids_of_last_fixtures_of_team {
		analyzed++
		fixture := fixture_map[id]
		if fixture.Goals.Home == fixture.Goals.Away {
			draws++
		} else if (fixture.Goals.Home > fixture.Goals.Away && team_id == fixture.Teams.Home.Id) || (fixture.Goals.Home < fixture.Goals.Away && team_id == fixture.Teams.Away.Id) {
			wins++
		} else {
			losses++
		}
	}
	return analyzed, wins, draws, losses
}

// Build Feature csv and save in local_datasets/prepared
func buildCsv(pre_game_stats_map map[int]FixturePreGameStats, fixture_map map[int]fixtures.FixtureItem, filename string) {

	// Create CSV
	p := filepath.Join("local_datasets", "prepared", filename)
	csvFile, err := os.Create(p)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	// Write header
	header := []string{"GameId", "HomeGamesAnalyzed", "HomeWins", "HomeDraws", "HomeLosses", "AwayGamesAnalyzed", "AwayWins", "AwayDraws", "AwayLosses", "Result"}
	csvwriter.Write(header)

	for key, item := range pre_game_stats_map {
		var result string
		fixture := fixture_map[key]

		// Don't include unplayed games
		if fixture.Meta.Status.Long == "Not Started" || fixture.Meta.Status.Long == "Time to be defined"{
			continue
		}

		if fixture.Goals.Home > fixture.Goals.Away {
			result = "Home"
		} else if fixture.Goals.Home == fixture.Goals.Away {
			result = "Draw"
		} else {
			result = "Away"
		}

		row := []string{
			fmt.Sprint(item.Id),
			fmt.Sprint(item.Home.GamesAnalzed),
			fmt.Sprint(item.Home.Wins),
			fmt.Sprint(item.Home.Draws),
			fmt.Sprint(item.Home.Losses),
			fmt.Sprint(item.Away.GamesAnalzed),
			fmt.Sprint(item.Away.Wins),
			fmt.Sprint(item.Away.Draws),
			fmt.Sprint(item.Away.Losses),
			result,
		}
		csvwriter.Write(row)
	}

	// Close
	csvwriter.Flush()
	csvFile.Close()
}

// Build Pregame Stat Map
func GetPregameStatsMap(fixture_list []fixtures.FixtureItem, fixture_map map[int]fixtures.FixtureItem) map[int]FixturePreGameStats {
	pre_game_stats_map := make(map[int]FixturePreGameStats)
	for i := 0; i < len(fixture_list); i++ {
		fixture := fixture_list[i]

		var pre_game_stat FixturePreGameStats
		pre_game_stat.Id = fixture.Meta.Id

		// Home Team Stats
		var home_stats TeamPreGameStats
		home_stats.LastXFixtureIds = getIdsOfLastXFixturesOfTeam(i, fixture.Teams.Home.Id, fixture_list, 10)
		home_stats.GamesAnalzed, home_stats.Wins, home_stats.Draws, home_stats.Losses = 
			getWdlStats(home_stats.LastXFixtureIds, fixture.Teams.Home.Id, fixture_map)
		pre_game_stat.Home = home_stats

		// Away Team Stats
		var away_stats TeamPreGameStats
		away_stats.LastXFixtureIds = getIdsOfLastXFixturesOfTeam(i, fixture.Teams.Away.Id, fixture_list, 10)
		away_stats.GamesAnalzed, away_stats.Wins, away_stats.Draws, away_stats.Losses = 
			getWdlStats(away_stats.LastXFixtureIds, fixture.Teams.Away.Id, fixture_map)
		pre_game_stat.Away = away_stats

		pre_game_stats_map[fixture.Meta.Id] = pre_game_stat
	}
	return pre_game_stats_map
}


func main() {
	// Get League ID
	id, err := leagues.GetLeagueId("Bundesliga", "Germany", true)
	if err != nil {
		log.Fatal(err)
	}

	// Get Fixtures as map (indexed by fixture id) and as list
	fixture_map, fixture_list := fixtures.GetFixtures(id, 2022, false)

	pre_game_stats_map := GetPregameStatsMap(fixture_list, fixture_map)

	buildCsv(pre_game_stats_map, fixture_map, "Bundesliga_2022.csv")
}