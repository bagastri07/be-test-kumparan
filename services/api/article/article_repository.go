package article

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type articleRepository struct{}

func (repo *articleRepository) GetTableName() string {
	return "articles"
}

func (repo *articleRepository) InsertArticle(ctx context.Context, tx *sqlx.Tx)
