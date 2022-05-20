package domain

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bjornnorgaard/laosyne/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/repository/database"
	"github.com/cockroachdb/errors"
)

func (a Api) GetPicture(ctx context.Context, filter string) (*model.Picture, error) {
	pictures, err := a.db.GetPicturesByFilter(ctx, database.GetPicturesByFilterParams{
		Column1: filter,
		Limit:   1,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get pic by filter: '%s'", filter)
	}

	if len(pictures) == 0 {
		return nil, errors.Newf("no pictures match filter: '%s'", filter)
	}

	pic := pictures[0]
	dto := &model.Picture{
		ID:        int(pic.ID),
		Path:      pic.Path,
		Ext:       pic.Ext,
		Views:     int(pic.Views),
		Likes:     int(pic.Likes),
		Rating:    pic.Rating,
		Deviation: pic.Deviation,
		Wins:      int(pic.Wins),
		Losses:    int(pic.Losses),
		Created:   pic.Created.String(),
		Updated:   pic.Updated.String(),
	}

	return dto, nil
}

func (a Api) AddPath(ctx context.Context, input model.NewPath) (*model.Path, error) {
	created, err := a.db.CreatePath(ctx, input.Path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create path")
	}

	dto := &model.Path{
		ID:      int(created.ID),
		Path:    created.Path,
		Created: created.Created.String(),
		Updated: created.Updated.String(),
	}

	return dto, nil
}

func (a Api) GetPaths(ctx context.Context) ([]*model.Path, error) {
	mediaPaths, err := a.db.GetPaths(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get media paths")
	}

	var dto []*model.Path
	for _, mp := range mediaPaths {
		dto = append(dto, &model.Path{
			ID:      int(mp.ID),
			Path:    mp.Path,
			Created: mp.Created.String(),
			Updated: mp.Updated.String(),
		})
	}

	return dto, nil
}

func (a Api) DeletePath(ctx context.Context, input model.DeletePath) (bool, error) {
	err := a.db.DeletePath(ctx, int64(input.PathID))
	if err != nil {
		return false, errors.Wrap(err, "failed to delete path")
	}
	return true, nil
}

func (a Api) ScanPath(ctx context.Context) (bool, error) {
	paths, err := a.db.GetPaths(ctx)
	if err != nil {
		return false, errors.Wrap(err, "failed to get paths")
	}

	go a.removeDeletedMedia()

	for _, p := range paths {
		go a.scanFolder(ctx, p.Path)
	}

	return true, nil
}

func (a Api) scanFolder(ctx context.Context, path string) {
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

		validExtension := contains(extensions, ext)
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

	a.gorm.InsertBulk(pictures)
}

func (a *Api) removeDeletedMedia() {
	var count int64
	batchSize := 100
	var batch []database.Picture

	a.gorm.Model(&database.Picture{}).Count(&count)

	var idsToDelete []int64
	for i := 0; i < int(count)/batchSize; i += 100 {
		a.gorm.Order("id").Offset(i).Limit(batchSize).Find(&batch)
		for _, image := range batch {
			_, err := os.Stat(image.Path)
			if err != nil {
				idsToDelete = append(idsToDelete, image.ID)
			}
		}
		batch = batch[:0]
	}

	if len(idsToDelete) == 0 {
		return
	}

	a.gorm.DeleteBulk(idsToDelete)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
