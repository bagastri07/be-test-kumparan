package article

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/jmoiron/sqlx"
)

type articleRepository struct{}

func NewRepository() ArticleRepository {
	return &articleRepository{}
}

var (
	searchFields = []string{
		"title",
		"body",
	}
)

func (repo *articleRepository) GetTableName() string {
	return "articles"
}

func (repo *articleRepository) buildInsertQuery(data *models.Article) sq.InsertBuilder {
	values := sq.Eq{
		"id":        data.ID,
		"author_id": data.AuthorID,
		"title":     data.Title,
		"body":      data.Body,
	}

	return sq.Insert(repo.GetTableName()).SetMap(values)
}

func (repo *articleRepository) InsertArticle(ctx context.Context, tx *sqlx.Tx, data *models.Article) error {
	segment := utils.StartTracer(ctx, "ArticleRepository", "InsertArticle")
	defer segment.End()

	query, args, err := repo.buildInsertQuery(data).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *articleRepository) buildSelectQuery() sq.SelectBuilder {
	fields := []string{
		repo.GetTableName() + ".id",
		repo.GetTableName() + ".author_id",
		repo.GetTableName() + ".title",
		repo.GetTableName() + ".body",
		repo.GetTableName() + ".created_at",
		repo.GetTableName() + ".updated_at",
		repo.GetTableName() + ".deleted_at",
		"authors.name AS author_name",
	}

	builder := sq.Select(fields...).
		Where(sq.Eq{
			repo.GetTableName() + ".deleted_at": nil,
		}).
		From(repo.GetTableName()).
		Join("authors ON authors.id = " + repo.GetTableName() + ".author_id")
	return builder

}

func (repo *articleRepository) buildSearchQuery(selectBuilder sq.SelectBuilder, keyword string) sq.SelectBuilder {
	var (
		search = "%" + keyword + "%"
		or     []sq.Sqlizer
	)

	for _, field := range searchFields {
		like := sq.Like{}
		like[field] = search
		or = append(or, like)
	}

	return selectBuilder.Where(sq.Or(or))
}

func (repo *articleRepository) GetCountArticle(ctx context.Context, db *sqlx.DB, filter *models.ArticleFilter) (uint64, error) {
	segment := utils.StartTracer(ctx, "ArticleRepository", "GetCountArticle")
	defer segment.End()

	var (
		count uint64
	)

	queryBuilder := repo.buildSelectQuery()

	if filter.Query != "" {
		queryBuilder = repo.buildSearchQuery(queryBuilder, filter.Query)
	}

	if filter.Author != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{
			"authors.name": filter.Author,
		})
	}

	query, args, err := sq.Select("count(*)").FromSelect(queryBuilder, "c").ToSql()
	if err != nil {
		return 0, err
	}

	if err := db.GetContext(ctx, &count, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return count, nil
}

func (repo *articleRepository) GetArticlesList(ctx context.Context, db *sqlx.DB, filter *models.ArticleFilter) ([]models.Article, error) {
	segment := utils.StartTracer(ctx, "ArticleRepository", "GetArticlesList")
	defer segment.End()

	articles := make([]models.Article, 0)

	queryBuilder := repo.buildSelectQuery().
		Limit(filter.Limit)

	if filter.Page > 1 {
		queryBuilder = queryBuilder.Offset((filter.Page - 1) * filter.Limit)

	}

	if filter.Query != "" {
		queryBuilder = repo.buildSearchQuery(queryBuilder, filter.Query)
	}

	if filter.Author != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{
			"authors.name": filter.Author,
		})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	if err := db.SelectContext(ctx, &articles, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return articles, nil
		}
		return nil, err
	}

	return articles, nil

}
