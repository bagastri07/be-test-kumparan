package models

type Author struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type CreateAuthorPayload struct {
	Name string `json:"name" validate:"required"`
}
