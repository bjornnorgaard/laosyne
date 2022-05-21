package domain

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bjornnorgaard/laosyne/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

func (a Api) GetPicture(_ context.Context, filter string) (*model.Picture, error) {
	var picture database.Picture
	a.db.Where("path LIKE ?", fmt.Sprintf("%%%s%%", filter)).Limit(1).First(&picture)

	if picture.ID == 0 {
		return nil, errors.Newf("no picture match filter: '%s'", filter)
	}

	dto := &model.Picture{
		ID:        int(picture.ID),
		Path:      picture.Path,
		Ext:       picture.Ext,
		Views:     picture.Views,
		Likes:     picture.Likes,
		Rating:    picture.Rating,
		Deviation: picture.Deviation,
		Wins:      picture.Wins,
		Losses:    picture.Losses,
		Created:   picture.CreatedAt.String(),
		Updated:   picture.UpdatedAt.String(),
	}

	return dto, nil
}

func (a Api) AddPath(_ context.Context, input model.NewPath) (*model.Path, error) {
	path := database.Path{Path: input.Path}
	a.db.Create(&path)
	dto := &model.Path{ID: int(path.ID), Path: path.Path, Created: path.CreatedAt.String()}
	return dto, nil
}

func (a Api) GetPaths(_ context.Context) ([]*model.Path, error) {
	var paths []database.Path
	a.db.Find(&paths)

	var dto []*model.Path
	for _, mp := range paths {
		dto = append(dto, &model.Path{ID: int(mp.ID), Path: mp.Path, Created: mp.CreatedAt.String()})
	}

	return dto, nil
}

func (a Api) DeletePath(_ context.Context, input model.DeletePath) (bool, error) {
	a.db.Unscoped().Delete(&database.Path{}, input.PathID)
	return true, nil
}

func (a Api) ScanPath(ctx context.Context) (bool, error) {
	var paths []database.Path
	a.db.Find(&paths)

	for _, p := range paths {
		go a.scanFolder(ctx, p.Path)
	}

	go a.removeDeletedMedia()

	return true, nil
}

func (a Api) scanFolder(ctx context.Context, path string) {
	//goland:noinspection ALL
	if runtime.GOOS != "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}

	_, err := os.Stat(path)
	if err != nil {
		return
	}

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	extensions := []string{".jpg", ".jpe", ".bmp", ".gif", ".png", ".webm"}
	var pictures []database.Picture
	for _, info := range dir {
		itemPath := strings.ToLower(filepath.Join(path, info.Name()))
		ext := filepath.Ext(itemPath)

		if info.IsDir() {
			a.scanFolder(ctx, itemPath)
			continue
		}

		validExtension := lo.Contains(extensions, ext)
		if !validExtension {
			continue
		}

		pictures = append(pictures, database.Picture{
			Path: itemPath,
			Ext:  ext,
		})
	}

	if len(pictures) == 0 {
		return
	}

	// a.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "path"}}, DoNothing: true}).Create(&pictures)
	a.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&pictures, 50)
}

func (a *Api) removeDeletedMedia() {
	var (
		limit        = 100
		offset       = 0
		pics         []database.Picture
		picsToDelete []database.Picture
	)

	for {
		a.db.Offset(offset).Limit(limit).Find(&pics)

		for _, p := range pics {
			_, err := os.Stat(p.Path)
			if err != nil {
				picsToDelete = append(picsToDelete, p)
			}
		}

		for _, p := range picsToDelete {
			a.db.Unscoped().Delete(&p)
		}

		if len(pics) < limit {
			break
		}

		offset += limit
		pics = pics[:0]
	}
}
