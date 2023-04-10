package db

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/tee8z/nullable"
)

type Result string

const (
	Home Result = "home"
	Away Result = "away"
	Draw Result = "draw"
	TBD  Result = "tbd"
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
	Id            uint      `gorm:"primary_key"`
	Time          time.Time `gorm:"not null; index"`
	HomeTeamId    uint      `gorm:"not null; index"`
	AwayTeamId    uint      `gorm:"not null; index"`
	HomeTeamGoals nullable.Int8
	AwayTeamGoals nullable.Int8
	Result        Result `sql:"type:result"`
	LeagueId      uint   `gorm:"not null; index"`
	League        League `gorm:"foreignKey:LeagueId"`
}
