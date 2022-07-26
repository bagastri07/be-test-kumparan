package article

import (
	"context"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/jmoiron/sqlx"
)

type ArticleRepository interface {
	GetTableName() string
	InsertArticle(ctx context.Context, tx *sqlx.Tx, data *models.Article) error
	GetCountArticle(ctx context.Context, db *sqlx.DB, filter *models.ArticleFilter) (uint64, error)
	GetArticlesList(ctx context.Context, db *sqlx.DB, filter *models.ArticleFilter) ([]models.Article, error)
}

type ArticleService interface {
	CreateArticle(ctx context.Context, payload *models.CreateArticlePayload) error
	GetArticlesPagination(ctx context.Context, filter *models.ArticleFilter) (*base_models.PaginationResponse, error)
}
