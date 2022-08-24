package ports

import (
	"contentSquare/src/internal/models"
	"context"
)

type DBRepository interface {
	IngestFileData(context.Context, string) error
	CountEvents(context.Context, models.Filters) (int64, error)
}
