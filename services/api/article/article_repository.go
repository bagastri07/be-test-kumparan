package article

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
)

type articleRepository struct{}

func NewRepository() ArticleRepository {
	return &articleRepository{}
}

func (repo *articleRepository) GetTableName() string {
	return "articles"
}

func (repo *articleRepository) buildInsertQuery(data *models.Article) sq.InsertBuilder {
	values := sq.Eq{
		"id":        data.ID,
		"author_id": data.AuthorID,
		"title":     data.Title,
		"body":      data.Body,
	}

	return sq.Insert(repo.GetTableName()).SetMap(values)
}

func (repo *articleRepository) InsertArticle(ctx context.Context, tx *sqlx.Tx, data *models.Article) error {
	query, args, err := repo.buildInsertQuery(data).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
