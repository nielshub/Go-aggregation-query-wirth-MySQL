package ports

import "context"

type DBRepository interface {
	IngestFileData(context.Context, string) error
}
