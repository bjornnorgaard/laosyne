package domain

import (
	"io/ioutil"
	"net/http"

	"github.com/bjornnorgaard/laosyne/graphql/graph/generated"
	"github.com/bjornnorgaard/laosyne/repository"
	"github.com/bjornnorgaard/laosyne/repository/database"
)

type Api struct {
	db repository.Repository
}

func (a Api) Mutation() generated.MutationResolver {
	return a
}

func (a Api) Query() generated.QueryResolver {
	return a
}

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

		fileBytes, err := ioutil.ReadFile(pic.Path)
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

func NewApi(r repository.Repository) *Api {
	return &Api{
		db: r,
	}
}
