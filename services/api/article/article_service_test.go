package article

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bagastri07/be-test-kumparan/mocks"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/models/base_models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_articleService_CreateArticle(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *models.CreateArticlePayload
	}

	type mockGetAuthorByID struct {
		data *models.Author
		err  error
	}

	type mockInsertArticle struct {
		err error
	}

	tests := []struct {
		name              string
		args              args
		mockGetAuthorByID mockGetAuthorByID
		mockInsertArticle mockInsertArticle
		wantErr           error
	}{
		{
			name: "when success, then return nil",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateArticlePayload{
					AuthorID: "123",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
				},
			},
			mockGetAuthorByID: mockGetAuthorByID{
				data: &models.Author{
					ID:   "123",
					Name: "Monkey D Luffy",
				},
				err: nil,
			},
			mockInsertArticle: mockInsertArticle{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "when author not found, then return err not found",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateArticlePayload{
					AuthorID: "123",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
				},
			},
			mockGetAuthorByID: mockGetAuthorByID{
				data: nil,
				err:  nil,
			},
			mockInsertArticle: mockInsertArticle{
				err: nil,
			},
			wantErr: utils.ErrNotFound,
		},
		{
			name: "when get author err, then return err",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateArticlePayload{
					AuthorID: "123",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
				},
			},
			mockGetAuthorByID: mockGetAuthorByID{
				data: nil,
				err:  errors.New("err"),
			},
			mockInsertArticle: mockInsertArticle{
				err: nil,
			},
			wantErr: errors.New("err"),
		},
		{
			name: "when insert article err, then return err",
			args: args{
				ctx: context.TODO(),
				payload: &models.CreateArticlePayload{
					AuthorID: "123",
					Title:    "Judul Kumparan",
					Body:     "Body Kumparan",
				},
			},
			mockGetAuthorByID: mockGetAuthorByID{
				data: &models.Author{
					ID:   "123",
					Name: "Monkey D Luffy",
				},
				err: nil,
			},
			mockInsertArticle: mockInsertArticle{
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
			authorRepository.On("GetAuthorByID", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockGetAuthorByID.data, tt.mockGetAuthorByID.err)

			articleRepository := new(mocks.ArticleRepository)
			articleRepository.On("InsertArticle", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockInsertArticle.err)

			svc := &articleService{
				db:                sqlxDB,
				articleRepository: articleRepository,
				authorReposiotry:  authorRepository,
			}

			err = svc.CreateArticle(tt.args.ctx, tt.args.payload)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_articleService_GetArticlesPagination(t *testing.T) {
	type args struct {
		ctx    context.Context
		filter *models.ArticleFilter
	}

	type mockGetArticlesList struct {
		data []models.Article
		err  error
	}

	type mockGetCountArticle struct {
		count uint64
		err   error
	}

	tests := []struct {
		name                string
		args                args
		mockGetArticlesList mockGetArticlesList
		mockGetCountArticle mockGetCountArticle
		want                *base_models.PaginationResponse
		wantErr             error
	}{
		{
			name: "when success, then return data",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockGetArticlesList: mockGetArticlesList{
				data: []models.Article{
					{
						ID:         "234",
						AuthorID:   "123",
						AuthorName: "Jojo",
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
			mockGetCountArticle: mockGetCountArticle{
				count: 1,
				err:   nil,
			},
			want: &base_models.PaginationResponse{
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
			wantErr: nil,
		},
		{
			name: "when get article list err, then return err",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockGetArticlesList: mockGetArticlesList{
				data: nil,
				err:  errors.New("err"),
			},
			mockGetCountArticle: mockGetCountArticle{
				count: 1,
				err:   nil,
			},
			want:    nil,
			wantErr: errors.New("err"),
		},
		{
			name: "when get count err, then return err",
			args: args{
				ctx:    context.TODO(),
				filter: &models.ArticleFilter{},
			},
			mockGetArticlesList: mockGetArticlesList{
				data: []models.Article{
					{
						ID:         "234",
						AuthorID:   "123",
						AuthorName: "Jojo",
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
			mockGetCountArticle: mockGetCountArticle{
				count: 0,
				err:   errors.New("err"),
			},
			want:    nil,
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

			articleRepository := new(mocks.ArticleRepository)
			articleRepository.On("GetArticlesList", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockGetArticlesList.data, tt.mockGetArticlesList.err)
			articleRepository.On("GetCountArticle", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockGetCountArticle.count, tt.mockGetCountArticle.err)

			svc := &articleService{
				db:                sqlxDB,
				articleRepository: articleRepository,
			}

			got, err := svc.GetArticlesPagination(tt.args.ctx, tt.args.filter)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
