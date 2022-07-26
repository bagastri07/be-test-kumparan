package author

import (
	"context"
	"database/sql"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
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
	segment := utils.StartTracer(ctx, "AuthorService", "CreateAuthor")
	defer segment.End()

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

	p := bluemonday.UGCPolicy()

	data := models.Author{
		ID:   utils.GenerateUUID(),
		Name: p.Sanitize(payload.Name),
	}

	if err := svc.authorRepository.InsertAuthor(ctx, tx, &data); err != nil {
		return err
	}

	return nil
}
