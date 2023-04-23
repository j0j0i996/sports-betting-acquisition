package sync

import (
	"db/client"
	"db/model"
	"fdu/acquisition"
	"fdu/transform"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm/clause"
)

var db_client = client.GetClient()

func formatSeason(season string) string {
	seasonWithoutSlash := strings.Replace(season, "/", "", -1)
	seasonNumber := seasonWithoutSlash[2:]
	return seasonNumber
}

func Sync(league model.League, season string) {
	seasonNumber := formatSeason(season)
	fmt.Println("Requesting data for " + league.Name + " in " + seasonNumber)
	acquired_fixtures, err := acquisition.GetHistoricData(seasonNumber, league.Name)
	if err != nil {
		log.Fatal("error in data acquisition" + err.Error())
	}
	fmt.Println("Received Data")
	for _, fixture := range acquired_fixtures {
		historic_odds, err := transform.HistoricOddsSourceModelToDbModel(fixture)
		if err != nil {
			log.Fatal("error in data acquisition" + err.Error())
		}
		db_client.Clauses(clause.OnConflict{DoNothing: true}).Create(&historic_odds)
	}

}
