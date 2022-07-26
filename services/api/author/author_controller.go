package author

import (
	"net/http"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/labstack/echo/v4"
)

type authorController struct {
	authorService AuthorService
}

func NewController(authorService AuthorService) *authorController {
	return &authorController{
		authorService: authorService,
	}
}

func (ctrl *authorController) HandleCreateAuthor(c echo.Context) error {
	segments := utils.StartControllerTracer(c, "AuthorController", "HandleCreateAuthor")
	defer segments.End()

	ctx := c.Request().Context()

	payload := new(models.CreateAuthorPayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctrl.authorService.CreateAuthor(ctx, payload); err != nil {
		return err
	}

	resp := models.MessageResponse{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
