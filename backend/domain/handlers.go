package domain

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (a Api) GetFile() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id := request.URL.Query()["id"]
		if len(id) != 1 {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var pic database.Picture
		a.db.First(&pic, id)

		if pic.ID == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		fileBytes, err := os.ReadFile(pic.Path)
		if err != nil {
			writer.WriteHeader(http.StatusExpectationFailed)
			return
		}

		_, err = writer.Write(fileBytes)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func (a Api) GetPicture(_ context.Context, input *model.SearchFilter) (*model.Picture, error) {
	var pic database.Picture
	a.buildQuery(input).Limit(1).First(&pic)

	if pic.ID == 0 {
		return nil, errors.New(fmt.Sprintf("no picture matches filter: '%s'", input.PathContains))
	}

	dto := &model.Picture{
		ID:        int(pic.ID),
		Path:      pic.Path,
		Ext:       pic.Ext,
		Views:     pic.Views,
		Likes:     pic.Likes,
		Rating:    pic.Rating,
		Deviation: pic.Deviation,
		Wins:      pic.Wins,
		Losses:    pic.Losses,
		CreatedAt: pic.CreatedAt.String(),
		UpdatedAt: pic.UpdatedAt.String(),
	}

	return dto, nil
}

func (a Api) buildQuery(input *model.SearchFilter) *gorm.DB {
	query := a.db.Session(&gorm.Session{})

	if input == nil {
		return query
	}

	if input.PathContains != nil {
		query = query.Where("path LIKE ?", fmt.Sprintf("%%%s%%", *input.PathContains))
	}

	return query
}

func (a Api) GetPictures(ctx context.Context, input *model.SearchFilter) ([]*model.Picture, error) {
	var pics []database.Picture
	a.buildQuery(input).Limit(100).Find(&pics)

	var dto []*model.Picture
	for _, p := range pics {
		dto = append(dto, &model.Picture{
			ID:        int(p.ID),
			Path:      p.Path,
			Ext:       p.Ext,
			Views:     p.Views,
			Likes:     p.Likes,
			Rating:    p.Rating,
			Deviation: p.Deviation,
			Wins:      p.Wins,
			Losses:    p.Losses,
			CreatedAt: p.CreatedAt.String(),
			UpdatedAt: p.UpdatedAt.String(),
		})
	}

	return dto, nil
}

func (a Api) AddPath(_ context.Context, input model.NewPath) (*model.Path, error) {
	path := database.Path{Path: input.Path}
	a.db.Create(&path)
	dto := &model.Path{ID: int(path.ID), Path: path.Path, CreatedAt: path.CreatedAt.String()}
	return dto, nil
}

func (a Api) GetPaths(_ context.Context) ([]*model.Path, error) {
	var paths []database.Path
	a.db.Find(&paths)

	var dto []*model.Path
	for _, mp := range paths {
		dto = append(dto, &model.Path{ID: int(mp.ID), Path: mp.Path, CreatedAt: mp.CreatedAt.String()})
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
