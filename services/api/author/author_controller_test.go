package author

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagastri07/be-test-kumparan/mocks"
	"github.com/bagastri07/be-test-kumparan/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_authorController_HandleCreateAuthor(t *testing.T) {
	type mockCreateAuthor struct {
		err error
	}

	tests := []struct {
		name             string
		payload          string
		mockCreateAuthor mockCreateAuthor
		wantErr          bool
		wantStatus       int
	}{
		{
			name: "when success, then return http.StatusOk",
			payload: `{
				"name": "Muhammad Iqbal"
			}`,
			mockCreateAuthor: mockCreateAuthor{
				err: nil,
			},
			wantErr:    false,
			wantStatus: http.StatusOK,
		},
		{
			name: "when validator fail, then return http.BadRequest",
			payload: `{
				
			}`,
			mockCreateAuthor: mockCreateAuthor{
				err: nil,
			},
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "when insert author err, then return http.InternalServerError",
			payload: `{
				"name": "Muhammad Iqbal"
			}`,
			mockCreateAuthor: mockCreateAuthor{
				err: errors.New("err"),
			},
			wantErr:    true,
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := new(mocks.AuthorService)
			svc.On("CreateAuthor", mock.Anything, mock.Anything).Return(tt.mockCreateAuthor.err)

			ctrl := NewController(svc)

			e := echo.New()

			validator.Init(e)

			req := httptest.NewRequest(http.MethodPost, "/authors", strings.NewReader(tt.payload))
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			c := e.NewContext(req, rec)
			c.SetPath("/authors")

			err := ctrl.HandleCreateAuthor(c)
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
