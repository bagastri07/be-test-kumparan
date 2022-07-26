package models

import "github.com/bagastri07/be-test-kumparan/models/base_models"

type Article struct {
	ID       string `db:"id"`
	AuthorID string `db:"author_id"`
	Title    string `db:"title"`
	Body     string `db:"body"`
	base_models.BaseTimestamp
}

type CreateArticlePayload struct {
	AuthorID string `json:"authorId" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
}
