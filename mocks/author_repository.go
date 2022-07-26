// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/bagastri07/be-test-kumparan/models"
	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"

	testing "testing"
)

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// GetAuthorByID provides a mock function with given fields: ctx, db, authorID
func (_m *AuthorRepository) GetAuthorByID(ctx context.Context, db *sqlx.DB, authorID string) (*models.Author, error) {
	ret := _m.Called(ctx, db, authorID)

	var r0 *models.Author
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.DB, string) *models.Author); ok {
		r0 = rf(ctx, db, authorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.DB, string) error); ok {
		r1 = rf(ctx, db, authorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTableName provides a mock function with given fields:
func (_m *AuthorRepository) GetTableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// InsertAuthor provides a mock function with given fields: ctx, tx, data
func (_m *AuthorRepository) InsertAuthor(ctx context.Context, tx *sqlx.Tx, data *models.Author) error {
	ret := _m.Called(ctx, tx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *models.Author) error); ok {
		r0 = rf(ctx, tx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthorRepository creates a new instance of AuthorRepository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthorRepository(t testing.TB) *AuthorRepository {
	mock := &AuthorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
