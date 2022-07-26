package models

import "github.com/bagastri07/be-test-kumparan/models/base_models"

type Author struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	base_models.BaseTimestamp
}

type CreateAuthorPayload struct {
	Name string `json:"name" validate:"required"`
}
