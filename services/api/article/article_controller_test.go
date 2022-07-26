package article

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bagastri07/be-test-kumparan/mocks"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/bagastri07/be-test-kumparan/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_articleController_HandleCreateArticle(t *testing.T) {
	type mockCreateArticle struct {
		err error
	}

	tests := []struct {
		name              string
		payload           string
		mockCreateArticle mockCreateArticle
		wantErr           bool
		wantStatus        int
	}{
		{
			name: "when success, then return http.Ok",
			payload: `{
				"authorId": "aee2ac32-ccd1-485b-b9ff-ba38de27ea52",
				"title": "Bola Merupakan objek",
				"body": "Berita ada di Kumparan"
			}`,
			mockCreateArticle: mockCreateArticle{
				err: nil,
			},
			wantErr:    false,
			wantStatus: http.StatusOK,
		},
		{
			name: "when validator fail, then return http.BadRequest",
			payload: `{
				"authorId": "aee2ac32-ccd1-485b-b9ff-ba38de27ea52",
				"title": "Bola Merupakan objek"
			}`,
			mockCreateArticle: mockCreateArticle{
				err: nil,
			},
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "when create article err, then return http.InternalServerError",
			payload: `{
				"authorId": "aee2ac32-ccd1-485b-b9ff-ba38de27ea52",
				"title": "Bola Merupakan objek",
				"body": "Berita ada di Kumparan"
			}`,
			mockCreateArticle: mockCreateArticle{
				err: errors.New("err"),
			},
			wantErr:    true,
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(mocks.ArticleService)
			svc.On("CreateArticle", mock.Anything, mock.Anything).Return(tt.mockCreateArticle.err)

			ctrl := NewController(svc)

			e := echo.New()

			validator.Init(e)

			req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(tt.payload))
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath("/articles")

			err := ctrl.HandleCreateArticle(c)
			if err != nil {
				//Mostly it's echo  behaviour like custom error handler and validator
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tt.wantStatus, he.Code)
				} else {
					assert.Equal(t, tt.wantErr, err != nil)
				}
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
			}
		})
	}
}

func Test_articleController_HandleGetArticlesPagination(t *testing.T) {
	type mockGetArticlesPagination struct {
		data *base_models.PaginationResponse
		err  error
	}

	tests := []struct {
		name                      string
		mockGetArticlesPagination mockGetArticlesPagination
		wantErr                   bool
		wantStatus                int
	}{
		{
			name: "when success, then return http.StatusOk",
			mockGetArticlesPagination: mockGetArticlesPagination{
				data: &base_models.PaginationResponse{
					Limit:      5,
					Page:       1,
					TotalItems: 1,
					TotalPages: 1,
					Items: []models.GetArticleResponse{
						{
							ID:         "234",
							AuthorName: "Jojo",
							Title:      "Judul Kumparan",
							Body:       "Body Kumparan",
							BaseTimeStampResponse: base_models.BaseTimeStampResponse{
								CreatedAt: &time.Time{},
								UpdatedAt: &time.Time{},
							},
						},
					},
				},
				err: nil,
			},
			wantErr:    false,
			wantStatus: http.StatusOK,
		},
		{
			name: "when get pagination err, then return http.InternalServerErr",
			mockGetArticlesPagination: mockGetArticlesPagination{
				data: nil,
				err:  errors.New("err"),
			},
			wantErr:    true,
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(mocks.ArticleService)
			svc.On("GetArticlesPagination", mock.Anything, mock.Anything).Return(tt.mockGetArticlesPagination.data, tt.mockGetArticlesPagination.err)

			ctrl := NewController(svc)

			e := echo.New()

			validator.Init(e)

			req := httptest.NewRequest(http.MethodGet, "/articles", nil)
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath("/articles")

			err := ctrl.HandleGetArticlesPagination(c)
			if err != nil {
				//Mostly it's echo  behaviour like custom error handler and validator
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tt.wantStatus, he.Code)
				} else {
					assert.Equal(t, tt.wantErr, err != nil)
				}
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
			}
		})
	}
}
