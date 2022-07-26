package models

import (
	"github.com/bagastri07/be-test-kumparan/constants"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/microcosm-cc/bluemonday"
)

type Article struct {
	ID         string `db:"id"`
	AuthorID   string `db:"author_id"`
	AuthorName string `db:"author_name"`
	Title      string `db:"title"`
	Body       string `db:"body"`
	base_models.BaseTimestamp
}

type CreateArticlePayload struct {
	AuthorID string `json:"authorId" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
}

type ArticleFilter struct {
	Query  string `query:"query"`
	Author string `query:"author"`
	base_models.PaginationParams
}

func (f *ArticleFilter) ProcessSanitize() {
	p := bluemonday.UGCPolicy()
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit > constants.MaxLimit {
		f.Limit = constants.MaxLimit
	}
	if f.Limit <= constants.MinLimit {
		f.Limit = constants.MinLimit
	}
	if f.Query != "" {
		f.Query = p.Sanitize(f.Query)
	}
	if f.Author != "" {
		f.Author = p.Sanitize(f.Author)

	}
}

type GetArticleResponse struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	base_models.BaseTimeStampResponse
}
