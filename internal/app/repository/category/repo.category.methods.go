package category

import (
	"context"
	"database/sql"
	"fmt"
	"siap_app/internal/app/entity/category"
	"strconv"

	"github.com/pkg/errors"
)

func (r *repository) GetCategoryAll(ctx context.Context, input category.PaginationCategory) ([]category.ResponseCategory, int64, error) {

	var (
		dataCategoryList []category.ResponseCategory
		offset           int
		query            string
		countQuery       string
		total            int64
	)

	offset = (input.Page - 1) * input.Size

	query = "SELECT " + columnSelectCategory + " FROM category WHERE 1=1"
	countQuery = "SELECT COUNT(*) FROM category WHERE 1=1"

	var args []interface{}
	var nextLimit int

	if input.Filter != nil {
		if input.Filter.CategoryName != "" {
			query += " AND category_name = $" + strconv.Itoa(len(args)+1)
			countQuery += " AND category_name = $" + strconv.Itoa(len(args)+1)
			args = append(args, input.Filter.CategoryName)
			nextLimit++
		}
		if input.Filter.Description != "" {
			query += " AND description LIKE $" + strconv.Itoa(len(args)+1)
			countQuery += " AND description LIKE $" + strconv.Itoa(len(args)+1)
			args = append(args, "%"+input.Filter.Description+"%")
			nextLimit++
		}
	}

	query += " ORDER BY category_id ASC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
	args = append(args, input.Size, offset)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dataCategory category.ResponseCategory
		if err := rows.Scan(
			&dataCategory.CategoryId,
			&dataCategory.CategoryName,
			&dataCategory.Description,
			&dataCategory.EntryUser,
			&dataCategory.EntryTime,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		dataCategoryList = append(dataCategoryList, dataCategory)
	}

	err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count items: %w", err)
	}

	return dataCategoryList, total, nil
}

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
