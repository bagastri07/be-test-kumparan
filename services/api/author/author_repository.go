package author

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
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
			"id":         authorID,
			"deleted_at": nil,
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	author := new(models.Author)

	if err := db.GetContext(ctx, author, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return author, nil
}
