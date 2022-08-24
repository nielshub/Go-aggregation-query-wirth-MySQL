package ports

import (
	"contentSquare/src/internal/models"
	"context"
)

type DBRepository interface {
	CountEvents(context.Context, models.Filters) (int64, error)
	CountDistinctUsers(context.Context, models.Filters) (int64, error)
	Exists(context.Context, models.Filters) (bool, error)
}
