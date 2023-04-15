package client

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "db/model"
)

func GetClient() *gorm.DB {
	dsn := os.Getenv("SPORTSBETTINGDB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.League{})
	db.AutoMigrate(&model.Team{})
	db.AutoMigrate(&model.Fixture{})

	return db

}
