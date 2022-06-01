package domain

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"github.com/samber/lo"
	"gorm.io/gorm/clause"
)

func (a API) ScanPaths(ctx context.Context) (bool, error) {
	var paths []database.Path
	a.db.Find(&paths)

	for _, p := range paths {
		go a.scanFolder(ctx, p.Path)
	}

	go a.removeDeletedMedia()

	return true, nil
}

func (a API) scanFolder(ctx context.Context, path string) {
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

func (a *API) removeDeletedMedia() {
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
