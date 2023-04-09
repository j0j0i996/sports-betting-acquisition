package interfaces

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Result string

const (
	Home Result = "home"
	Away Result = "away"
	Draw Result = "draw"
)

func (r *Result) Scan(value interface{}) error {
	if value == nil {
		return errors.New("status: scan value is nil")
	}
	switch v := value.(type) {
	case []byte:
		*r = Result(v)
		return nil
	case string:
		*r = Result(v)
		return nil
	default:
		return errors.New("Result: Scan value is not []byte or string")
	}
}

func (r Result) Value() (driver.Value, error) {
	return string(r), nil
}

type Fixture struct {
	Id            uint      `gorm:"primary_key`
	Time          time.Time `gorm:"not null; index`
	HomeTeamId    uint      `gorm:"not null; index`
	AwayTeamId    uint      `gorm:"not null; index`
	HomeTeamGoals uint
	AwayTeamGoals uint
	Result        Result `sql:"type:result"`
}

func GetClient() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5434"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Fixture{})

	return db

}
