package article

import (
	"context"
	"net/http"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/labstack/echo/v4"
)

type articleController struct {
	articleService ArticleService
}

func NewController(articleService ArticleService) *articleController {
	return &articleController{
		articleService: articleService,
	}
}

func (ctrl *articleController) HandleCreateArticle(c echo.Context) error {
	payload := new(models.CreateArticlePayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctrl.articleService.CreateArticle(context.Background(), payload); err != nil {
		return err
	}

	resp := models.MessageResponse{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
