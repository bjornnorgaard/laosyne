package database

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"gorm.io/gorm"
)

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

func (p *Picture) SetPath(path string) {
	path = strings.ToLower(path)

	//goland:noinspection GoBoolExpressions
	if runtime.GOOS != "windows" {
		path = fmt.Sprintf("c:%s", path)
		path = strings.Replace(path, "/", "\\", -1)
	}
	p.Path = path
}

func (p Picture) GetPath() string {
	path := strings.ToLower(p.Path)

	//goland:noinspection GoBoolExpressions
	if runtime.GOOS != "windows" {
		path = strings.Replace(path, "c:", "", 1)
		path = strings.Replace(path, "\\", "/", -1)
	}

	return path
}

func (p Picture) ReadFile() ([]byte, error) {
	path := p.GetPath()

	return os.ReadFile(path)
}
