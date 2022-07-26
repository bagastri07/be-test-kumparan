package author

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
)

type authorRepository struct{}

func NewRepository() AuthorRepository {
	return &authorRepository{}
}

var (
	fields = []string{
		"id",
	}
)

func (repo *authorRepository) GetTableName() string {
	return "authors"
}

func (repo *authorRepository) buildInsertQuery(data *models.Author) sq.InsertBuilder {
	values := sq.Eq{
		"id":   data.ID,
		"name": data.Name,
	}

	return sq.Insert(repo.GetTableName()).SetMap(values)
}

func (repo *authorRepository) InsertAuthor(ctx context.Context, tx *sqlx.Tx, data *models.Author) error {

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

func (repo *authorRepository) GetAuthorByID(ctx context.Context, db *sqlx.DB, authorID string) (*models.Author, error) {
	query, args, err := sq.Select(fields...).
		From(repo.GetTableName()).
		Where(sq.Eq{
			"author_id":  authorID,
			"deleted_at": nil,
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	author := models.Author{}

	if err := db.SelectContext(ctx, &author, query, args...); err != nil {
		return nil, err
	}

	return &author, nil
}
