package article

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/services/api/author"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
)

type articleSercive struct {
	db                *sqlx.DB
	articleRepository ArticleRepository
	authorReposiotry  author.AuthorRepository
}

func NewService(db *sqlx.DB, artcarticleRepository ArticleRepository, authorRepository author.AuthorRepository) ArticleService {
	return &articleSercive{
		db:                db,
		articleRepository: artcarticleRepository,
		authorReposiotry:  authorRepository,
	}
}

func (svc *articleSercive) CreateArticle(ctx context.Context, payload *models.CreateArticlePayload) error {

	author, err := svc.authorReposiotry.GetAuthorByID(ctx, svc.db, payload.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return &utils.CustomError{
			Code:    http.StatusNotFound,
			Message: "author not found.",
		}
	}

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
