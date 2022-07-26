package article

import (
	"context"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
)

type ArticleRepository interface {
	GetTableName() string
	InsertArticle(ctx context.Context, tx *sqlx.Tx, data *models.Article) error
}
