package category

import (
	"context"
	"fmt"
	"siap_app/internal/app/entity/category"
)

func (uc *UseCase) CreateCategory(ctx context.Context, input category.RequestCategory) error {
	countDuplikat, err := uc.categoryRepo.FindByName(ctx, input.CategoryName)

	if err != nil {
		return fmt.Errorf("failled get category name")
	}

	if countDuplikat > 0 {
		return fmt.Errorf("category has been registered")
	}

	return uc.categoryRepo.CreateCategory(ctx, input)
}

func (uc *UseCase) GetCategoryById(ctx context.Context, ID int) (category.ResponseCategory, error) {
	data := category.ResponseCategory{}
	getCategoryResponse, err := uc.categoryRepo.GetCategoryById(ctx, ID)

	if err != nil {
		return data, err
	}

	return getCategoryResponse, nil
}

func (uc *UseCase) UpdateCategory(ctx context.Context, input category.RequestCategoryUpdate) error {

	return uc.categoryRepo.UpdateCategory(ctx, input)
}

func (uc *UseCase) DeleteCategory(ctx context.Context, id int) error {
	getCategoryResponse, err := uc.categoryRepo.GetCategoryById(ctx, id)

	if err != nil {
		return err
	}

	if getCategoryResponse.CategoryId == "" {
		return fmt.Errorf("fetch category failed")
	}

	return uc.categoryRepo.DeleteCategory(ctx, id)

}
