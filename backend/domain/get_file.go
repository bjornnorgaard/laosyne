package domain

import (
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) GetFile() http.Handler {
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

		pic.Views++
		a.db.Save(&pic)

		//goland:noinspection ALL
		if runtime.GOOS != "windows" {
			pic.Path = strings.Replace(pic.Path, "C:", "", 1)
			pic.Path = strings.Replace(pic.Path, "\\", "/", -1)
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
