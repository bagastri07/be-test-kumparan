package author

import (
	"context"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
)

type AuthorRepository interface {
	GetTableName() string
	InsertAuthor(ctx context.Context, tx *sqlx.Tx, data *models.Author) error
}

type AuthorService interface {
	CreateAuthor(ctx context.Context, payload *models.CreateAuthorPayload) error
}
