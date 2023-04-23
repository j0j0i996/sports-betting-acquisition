package insert

import (
	"db/client"
	"db/model"

	"gorm.io/gorm/clause"
)

var db_client = client.GetClient()

func InsertBookmaker(name string, slug string) {
	// Insert bookmakers
	bookmakers := [1]model.Bookmaker{
		{Name: name, Slug: slug},
	}
	db_client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(&bookmakers)
}
