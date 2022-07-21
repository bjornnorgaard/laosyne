package database

import (
	"gorm.io/gorm"
)

type Path struct {
	gorm.Model
	Path string `gorm:"uniqueIndex"`
}
