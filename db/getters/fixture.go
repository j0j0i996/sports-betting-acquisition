package getters

import (
	client "db/client"
	model "db/model"
	"errors"
	"time"
)

var db_client = client.GetClient()

// Get teams from db and cache
func FindFixtureFromTeamIdsAndTimestampInDB(input_time time.Time, home model.Team, away model.Team) (model.Fixture, error) {
	var fixtures []model.Fixture
	db_client.Model(&model.Fixture{HomeTeamId: home.Id, AwayTeamId: away.Id}).Find(&fixtures)

	// find fixture with closest timestamp
	min_difference, _ := time.ParseDuration("96h")
	closest_fixture := model.Fixture{}
	for _, fixture := range fixtures {
		diff := time.Duration.Abs(input_time.Sub(fixture.Time))
		if diff < min_difference {
			min_difference = diff
			closest_fixture = fixture
		}
	}
	if (closest_fixture == model.Fixture{}) {
		return closest_fixture, errors.New("no close fixture found")
	} else {
		return closest_fixture, nil
	}
}
