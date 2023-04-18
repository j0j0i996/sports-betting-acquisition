package getters

import (
	model "db/model"
)

// TODO: BUILD CACHE
func GetBookmakersBySlugInDB(slug string) model.Bookmaker {
	var bookmaker model.Bookmaker
	db_client.Model(&model.Bookmaker{Slug: slug}).First(&bookmaker)
	return bookmaker
}
