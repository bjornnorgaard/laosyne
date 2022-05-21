package database

import (
	"gorm.io/gorm"
)

type Path struct {
	gorm.Model
	Path string `gorm:"uniqueIndex"`
}

type Picture struct {
	gorm.Model
	Path      string `gorm:"uniqueIndex"`
	Ext       string
	Views     int
	Likes     int
	Rating    float64 `gorm:"index"`
	Deviation float64
	Wins      int
	Losses    int
}
