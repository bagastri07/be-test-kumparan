package author

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_authorRepository_InsertAuthor(t *testing.T) {
	type mockExec struct {
		err error
	}

	type args struct {
		ctx  context.Context
		data *models.Author
	}
	tests := []struct {
		name     string
		mockExec mockExec
		args     args
		wantErr  error
	}{
		{
			name: "when success, then return nil",
			args: args{
				ctx: context.TODO(),
				data: &models.Author{
					ID:   "1234",
					Name: "Boba",
				},
			},
			mockExec: mockExec{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "when err, then return err",
			args: args{
				ctx: context.TODO(),
				data: &models.Author{
					ID:   "1234",
					Name: "Boba",
				},
			},
			mockExec: mockExec{
				err: sql.ErrConnDone,
			},
			wantErr: sql.ErrConnDone,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewRepository()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			sqlxDB := sqlx.NewDb(db, "sqlmock")

			mock.ExpectBegin()
			tx, _ := sqlxDB.BeginTxx(tt.args.ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})

			mockExpectExec := mock.ExpectExec(`INSERT INTO authors (id,name) VALUES (?,?)`)

			if tt.mockExec.err != nil {
				mockExpectExec.WillReturnError(tt.mockExec.err)
			} else {
				mockExpectExec.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			err = repo.InsertAuthor(tt.args.ctx, tx, tt.args.data)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
