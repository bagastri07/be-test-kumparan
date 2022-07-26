package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCustomErr(code int, message string) *echo.HTTPError {

	return echo.NewHTTPError(code, CustomError{
		Code:    code,
		Message: message,
	})
}

// general errors
var (
	ErrNotFound = NewCustomErr(http.StatusNotFound, "not found.")
)
