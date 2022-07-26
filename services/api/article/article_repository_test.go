package article

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_articleRepository_InsertArticle(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *models.Article
	}

	type mockExec struct {
		err error
	}

	tests := []struct {
		name     string
		args     args
		mockExec mockExec
		wantErr  error
	}{
		{
			name: "when success, then return nil",
			args: args{
				ctx: context.TODO(),
				data: &models.Article{
					ID:       "1234",
					AuthorID: "2345",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
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
				data: &models.Article{
					ID:       "1234",
					AuthorID: "2345",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
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

			mockExpectExec := mock.ExpectExec(`INSERT INTO articles (author_id,body,id,title) VALUES (?,?,?,?)`)

			if tt.mockExec.err != nil {
				mockExpectExec.WillReturnError(tt.mockExec.err)
			} else {
				mockExpectExec.WillReturnResult(sqlmock.NewResult(1, 1))
			}

			err = repo.InsertArticle(tt.args.ctx, tx, tt.args.data)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_articleRepository_GetCountArticle(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter *models.ArticleFilter
	}

	type mockCount struct {
		count uint64
		err   error
	}

	tests := []struct {
		name      string
		args      args
		mockCount mockCount
		want      uint64
		wantErr   error
	}{
		{
			name: "when success, then return count",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockCount: mockCount{
				count: 34,
				err:   nil,
			},
			want:    34,
			wantErr: nil,
		},
		{
			name: "when err, then return err",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockCount: mockCount{
				count: 0,
				err:   sql.ErrConnDone,
			},
			want:    0,
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

			mockExpectGet := mock.ExpectQuery(
				`SELECT count(*) 
				FROM
					(SELECT
						articles.id,
						articles.author_id,
						articles.title,
						articles.body,
						articles.created_at,
						articles.updated_at,
						articles.deleted_at,
						authors.name AS author_name
					FROM
						articles
					JOIN
						authors ON authors.id = articles.author_id
					WHERE
						articles.deleted_at IS NULL
					ORDER BY created_at DESC) AS c`,
			)

			if tt.mockCount.err != nil {
				mockExpectGet.WillReturnError(tt.mockCount.err)
			} else {
				row := sqlmock.NewRows([]string{
					"count(*)",
				})

				row.AddRow(tt.mockCount.count)
				mockExpectGet.WillReturnRows(row)
			}

			data, err := repo.GetCountArticle(tt.args.ctx, sqlxDB, tt.args.filter)
			assert.Equal(t, tt.want, data)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_articleRepository_GetArticlesList(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter *models.ArticleFilter
	}

	type mockSelect struct {
		data []models.Article
		err  error
	}

	tests := []struct {
		name       string
		args       args
		mockSelect mockSelect
		want       []models.Article
		wantErr    error
	}{
		{
			name: "when success, then return data",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockSelect: mockSelect{
				data: []models.Article{
					{
						ID:         "1234",
						AuthorID:   "2345",
						AuthorName: "Jamal",
						Title:      "Judul Kumparan",
						Body:       "Body Kumparan",
						BaseTimestamp: base_models.BaseTimestamp{
							CreatedAt: &time.Time{},
							UpdatedAt: &time.Time{},
						},
					},
				},
				err: nil,
			},
			want: []models.Article{
				{
					ID:         "1234",
					AuthorID:   "2345",
					AuthorName: "Jamal",
					Title:      "Judul Kumparan",
					Body:       "Body Kumparan",
					BaseTimestamp: base_models.BaseTimestamp{
						CreatedAt: &time.Time{},
						UpdatedAt: &time.Time{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "when err, then return err",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockSelect: mockSelect{
				data: nil,
				err:  sql.ErrConnDone,
			},
			want:    nil,
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

			mockExpectSelect := mock.ExpectQuery(
				`SELECT
					articles.id,
					articles.author_id,
					articles.title,
					articles.body,
					articles.created_at,
					articles.updated_at,
					articles.deleted_at,
					authors.name AS author_name
				FROM
					articles
				JOIN
					authors ON authors.id = articles.author_id
				WHERE articles.deleted_at IS NULL ORDER BY created_at DESC LIMIT 0`,
			)

			if tt.mockSelect.err != nil {
				mockExpectSelect.WillReturnError(tt.mockSelect.err)
			} else {
				rows := sqlmock.NewRows([]string{
					"id",
					"author_id",
					"author_name",
					"title",
					"body",
					"created_at",
					"updated_at",
				})
				for _, article := range tt.mockSelect.data {
					rows.AddRow(
						article.ID,
						article.AuthorID,
						article.AuthorName,
						article.Title,
						article.Body,
						article.CreatedAt,
						article.UpdatedAt,
					)
				}
				mockExpectSelect.WillReturnRows(rows)
			}

			data, err := repo.GetArticlesList(tt.args.ctx, sqlxDB, tt.args.filter)
			assert.Equal(t, tt.want, data)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
