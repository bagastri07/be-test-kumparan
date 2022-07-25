package author

import (
	"context"
	"database/sql"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
)

type authorService struct {
	db               *sqlx.DB
	authorRepository AuthorRepository
}

func NewService(db *sqlx.DB, authorRepository AuthorRepository) AuthorService {
	return &authorService{
		db:               db,
		authorRepository: authorRepository,
	}
}

func (svc *authorService) CreateAuthor(ctx context.Context, payload *models.CreateAuthorPayload) error {
	tx, err := svc.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	data := models.Author{
		ID:   utils.GenerateUUID(),
		Name: payload.Name,
	}

	if err := svc.authorRepository.InsertAuthor(ctx, tx, &data); err != nil {
		return err
	}

	return nil
}
