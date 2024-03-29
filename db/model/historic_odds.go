package model

type HistoricOdds struct {
	Id          uint      `gorm:"primary_key"`
	FixtureId   uint      `gorm:"not null; index:idx_member,priority:1"`
	Fixture     Fixture   `gorm:"foreignKey:FixtureId"`
	BookmakerId uint      `gorm:"not null; index:idx_member,priority:2"`
	Bookmaker   Bookmaker `gorm:"foreignKey:BookmakerId"`
	HomeOdds    float64   `gorm:"not null"`
	DrawOdds    float64   `gorm:"not null"`
	AwayOdds    float64   `gorm:"not null"`
}
