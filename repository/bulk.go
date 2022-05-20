package repository

import (
	"database/sql"
	"log"

	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormContext struct {
	*gorm.DB
}

func NewGormContext(db *sql.DB) GormContext {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	if err != nil {
		log.Fatal(errors.Wrap(err, "GORM failed to reuse existing connection"))
	}

	return GormContext{gormDB}
}

func (g GormContext) InsertBulk(pictures []database.Picture) {
	g.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "path"}},
		DoNothing: true,
	}).Create(&pictures)
}

func (g GormContext) DeleteBulk(pictureIds []int64) {
	g.Delete(database.Picture{}, pictureIds)
}
