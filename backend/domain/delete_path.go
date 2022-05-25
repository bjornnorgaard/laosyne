package domain

import (
	"context"

	"github.com/bjornnorgaard/laosyne/backend/graphql/graph/model"
	"github.com/bjornnorgaard/laosyne/backend/repository/database"
)

func (a Api) DeletePath(_ context.Context, input model.DeletePath) (bool, error) {
	a.db.Unscoped().Delete(&database.Path{}, input.PathID)
	return true, nil
}
