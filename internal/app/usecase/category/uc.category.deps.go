package category

import (
	"context"
	"siap_app/internal/app/entity/category"
)

type categoryRepo interface {
	GetCategoryAll(ctx context.Context, input category.PaginationCategory) ([]category.ResponseCategory, int64, error)
	CreateCategory(ctx context.Context, input category.RequestCategory) error
	GetCategoryById(ctx context.Context, id int) (category.ResponseCategory, error)
	DeleteCategory(ctx context.Context, id int) error
	UpdateCategory(ctx context.Context, input category.RequestCategoryUpdate) error
	FindByName(ctx context.Context, name string) (int, error)
}
