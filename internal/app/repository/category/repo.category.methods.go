package category

import (
	"context"
	"database/sql"
	"siap_app/internal/app/entity/category"

	"github.com/pkg/errors"
)

func (r *repository) CreateCategory(ctx context.Context, input category.RequestCategory) error {
	_, err := r.db.ExecContext(ctx, queryInsertCategory,
		input.CategoryName,
		input.Description,
		input.EntryUser,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create category")
	}

	return nil
}

func (r *repository) GetCategoryById(ctx context.Context, category_id int) (category.ResponseCategory, error) {
	var category category.ResponseCategory
	err := r.db.QueryRowContext(ctx, queryCategoryById, category_id).Scan(
		&category.CategoryId,
		&category.CategoryName,
		&category.Description,
		&category.EntryUser,
		&category.EntryTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return category, errors.Wrap(err, "category not found")
		}
		return category, errors.Wrap(err, "failed to get category by id")
	}

	return category, nil
}

func (r *repository) UpdateCategory(ctx context.Context, input category.RequestCategoryUpdate) error {
	_, err := r.db.ExecContext(ctx, queryUpdateCategory,
		input.CategoryName,
		input.Description,
		input.ID,
	)

	if err != nil {
		return errors.Wrap(err, "failed to update category")
	}

	return nil
}

func (r *repository) DeleteCategory(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, queryDeleteCategory,
		id,
	)

	if err != nil {
		return errors.Wrap(err, "failed to delete category")
	}

	return nil
}

func (r *repository) FindByName(ctx context.Context, name string) (int, error) {
	var countData int
	err := r.db.QueryRowContext(ctx, queryCategoryByName, name).Scan(
		&countData,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return countData, errors.Wrap(err, "category not found")
		}
		return countData, errors.Wrap(err, "failed to get category by name")
	}

	return countData, nil
}
