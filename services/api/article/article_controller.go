package article

import (
	"net/http"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
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
	segments := utils.StartControllerTracer(c, "ArticleController", "HandleCreateArticle")
	defer segments.End()

	ctx := c.Request().Context()

	payload := new(models.CreateArticlePayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctrl.articleService.CreateArticle(ctx, payload); err != nil {
		return err
	}

	resp := models.MessageResponse{
		Message: "OK",
	}

	return c.JSON(http.StatusOK, resp)
}

func (ctrl *articleController) HandleGetArticlesPagination(c echo.Context) error {
	segments := utils.StartControllerTracer(c, "ArticleController", "HandleGetArticlesPagination")
	defer segments.End()

	ctx := c.Request().Context()

	filter := new(models.ArticleFilter)

	if err := c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := ctrl.articleService.GetArticlesPagination(ctx, filter)
	if err != nil {
		return err
	}

	resp := models.DataResponse{
		Data: result,
	}

	return c.JSON(http.StatusOK, resp)
}
