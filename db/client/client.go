package client

import (
	model "db/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClient() *gorm.DB {
	dsn := os.Getenv("SPORTSBETTINGDB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&model.League{})
	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Fixture{})
	db.AutoMigrate(&model.Bookmaker{})
	db.AutoMigrate(&model.HistoricOdds{})

	return db

}
