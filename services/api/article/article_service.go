package article

import (
	"context"
	"database/sql"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
)

type articleSercive struct {
	db                *sqlx.DB
	articleRepository ArticleRepository
}

func NewService(db *sqlx.DB, artcarticleRepository ArticleRepository) ArticleService {
	return &articleSercive{
		db:                db,
		articleRepository: artcarticleRepository,
	}
}

func (svc *articleSercive) CreateArticle(ctx context.Context, payload *models.CreateArticlePayload) error {
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

	data := models.Article{
		ID:       utils.GenerateUUID(),
		AuthorID: payload.AuthorID,
		Title:    payload.Title,
		Body:     payload.Body,
	}

	if err := svc.articleRepository.InsertArticle(ctx, tx, &data); err != nil {
		return err
	}

	return nil
}
