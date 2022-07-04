package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a API) DeletePath(_ context.Context, pathID int) (bool, error) {
	a.db.Unscoped().Delete(&database.Path{}, pathID)
	return true, nil
}
