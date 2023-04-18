package getters

import (
	model "db/model"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antzucaro/matchr"
)

// Get teams from db and cache
var cache_teams []model.Team

func GetTeamsInDB() []model.Team {
	if len(cache_teams) == 0 {
		fmt.Println("Get Data from db")
		db_client.Model(&model.Team{}).Find(&cache_teams)
	}
	return cache_teams
}

// Find team in db with most similar team name to input string
func GetTeamFromSimilarTeamName(similar_team_name string) (model.Team, error) {
	teams := GetTeamsInDB()
	min_str_distance := 3
	most_similar_team := model.Team{}
	for _, team := range teams {
		str_distance := matchr.LongestCommonSubsequence(
			strings.ToLower(team.Name),
			strings.ToLower(similar_team_name),
		)
		if str_distance > min_str_distance {
			most_similar_team = team
			min_str_distance = str_distance
		}
	}
	fmt.Println("Most similar team name for: " + similar_team_name + " -> " + most_similar_team.Name +
		". The score is: " + strconv.Itoa(min_str_distance))
	if (most_similar_team == model.Team{}) {
		return most_similar_team, errors.New("no similar team name found")
	} else {
		return most_similar_team, nil
	}
}
