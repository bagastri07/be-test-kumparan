package author

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/be-test-kumparan/mocks"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_authorService_CreateAuthor(t *testing.T) {
	type mockInsertAuthor struct {
		err error
	}

	type args struct {
		ctx     context.Context
		payload *models.CreateAuthorPayload
	}
	tests := []struct {
		name             string
		args             args
		mockInsertAuthor mockInsertAuthor
		wantErr          error
	}{
		{
			name: "when success, then return nil",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateAuthorPayload{
					Name: "Bambang",
				},
			},
			mockInsertAuthor: mockInsertAuthor{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "when insert author err, then return err",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateAuthorPayload{
					Name: "Bambang",
				},
			},
			mockInsertAuthor: mockInsertAuthor{
				err: errors.New("err"),
			},
			wantErr: errors.New("err"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			sqlxDB := sqlx.NewDb(db, "sqlMock")
			mockDB.ExpectBegin()

			authorRepository := new(mocks.AuthorRepository)
			authorRepository.On("InsertAuthor", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockInsertAuthor.err)

			svc := &authorService{
				db:               sqlxDB,
				authorRepository: authorRepository,
			}

			err = svc.CreateAuthor(tt.args.ctx, tt.args.payload)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
