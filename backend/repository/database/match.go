package database

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	Quality  float64
	WinnerID uint
	Winner   Picture
	LoserID  uint
	Loser    Picture
}
