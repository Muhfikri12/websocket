package categoryrepositpry_test

import (
	"fmt"
	"project/helper"
	categoryrepositpry "project/repository/category_repositpry"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestShowAllCategory(t *testing.T) {
	db, mock := helper.SetupTestDB()
	defer func() { _ = mock.ExpectationsWereMet() }()

	log := *zap.NewNop()
	categoryRepo := categoryrepositpry.NewCategoryRepo(db, &log)

	t.Run("Successfully show all categories", func(t *testing.T) {
		page, limit := 1, 2

		// Mock count query
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "categories" WHERE "categories"."deleted_at" IS NULL`)).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

		// Mock data query with pagination
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "categories" WHERE "categories"."deleted_at" IS NULL LIMIT $1`)).
			WithArgs(2).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Category 1").
				AddRow(2, "Category 2"))

		categories, totalCount, totalPages, err := categoryRepo.ShowAllCategory(page, limit)

		assert.NoError(t, err)
		assert.NotNil(t, categories)
		assert.Len(t, *categories, 2)

		assert.Equal(t, "Category 1", (*categories)[0].Name)
		assert.Equal(t, "Category 2", (*categories)[1].Name)

		assert.Equal(t, 5, totalCount)
		assert.Equal(t, 3, totalPages) // 5 categories / 2 per page = 3 pages
	})

	t.Run("Failed to show all categories due to database error", func(t *testing.T) {
		page, limit := 1, 2

		// Mock count query
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "categories" WHERE "categories"."deleted_at" IS NULL`)).
			WillReturnError(fmt.Errorf("database error"))

		// Call the repository method
		categories, totalCount, totalPages, err := categoryRepo.ShowAllCategory(page, limit)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, categories)
		assert.Equal(t, 0, totalCount)
		assert.Equal(t, 0, totalPages)
		assert.EqualError(t, err, "database error")
	})
}
