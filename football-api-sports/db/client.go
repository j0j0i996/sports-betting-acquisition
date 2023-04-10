package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClient() *gorm.DB {
	dsn := os.Getenv("SPORTSBETTINGDB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&League{})
	db.AutoMigrate(&Fixture{})

	return db

}
