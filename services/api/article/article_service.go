package article

import (
	"context"
	"database/sql"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/bagastri07/be-test-kumparan/services/api/author"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
)

type articleService struct {
	db                *sqlx.DB
	articleRepository ArticleRepository
	authorReposiotry  author.AuthorRepository
}

func NewService(db *sqlx.DB, artcarticleRepository ArticleRepository, authorRepository author.AuthorRepository) ArticleService {
	return &articleService{
		db:                db,
		articleRepository: artcarticleRepository,
		authorReposiotry:  authorRepository,
	}
}

func (svc *articleService) CreateArticle(ctx context.Context, payload *models.CreateArticlePayload) error {
	segment := utils.StartTracer(ctx, "ArticleService", "CreateArticle")
	defer segment.End()

	author, err := svc.authorReposiotry.GetAuthorByID(ctx, svc.db, payload.AuthorID)
	if err != nil {
		return err
	}

	if author == nil {
		return utils.ErrNotFound
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

	p := bluemonday.UGCPolicy()

	data := models.Article{
		ID:       utils.GenerateUUID(),
		AuthorID: payload.AuthorID,
		Title:    p.Sanitize(payload.Title),
		Body:     p.Sanitize(payload.Body),
	}

	if err := svc.articleRepository.InsertArticle(ctx, tx, &data); err != nil {
		return err
	}

	return nil
}

func (svc *articleService) GetArticlesPagination(ctx context.Context, filter *models.ArticleFilter) (*base_models.PaginationResponse, error) {
	segment := utils.StartTracer(ctx, "ArticleService", "GetArticlesPagination")
	defer segment.End()

	filter.ProcessSanitize()

	articles, err := svc.articleRepository.GetArticlesList(ctx, svc.db, filter)
	if err != nil {
		return nil, err
	}

	totalItems, err := svc.articleRepository.GetCountArticle(ctx, svc.db, filter)
	if err != nil {
		return nil, err
	}

	items := make([]models.GetArticleResponse, 0)

	for _, article := range articles {

		items = append(items, models.GetArticleResponse{
			ID:         article.ID,
			AuthorName: article.AuthorName,
			Title:      article.Title,
			Body:       article.Body,
			BaseTimeStampResponse: base_models.BaseTimeStampResponse{
				CreatedAt: article.CreatedAt,
				UpdatedAt: article.UpdatedAt,
			},
		})
	}

	resp := base_models.PaginationResponse{}
	resp.ToResponse(&filter.PaginationParams, items, totalItems)

	return &resp, nil
}
