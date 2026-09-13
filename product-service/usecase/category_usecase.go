package usecase

import (
	"context"
	"pos/product-service/model"
	"pos/product-service/repository"
)

type CategoryUsecaseInterface interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetAllCategories(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.Category, int64, error)
	GetCategoryByID(ctx context.Context, id uint) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) error
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepositoryInterface
}

func (c *categoryUsecase) CreateCategory(ctx context.Context, category *model.Category) error {
	return c.categoryRepo.CreateCategory(ctx, category)
}

func (c *categoryUsecase) GetAllCategories(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.Category, int64, error) {
	return c.categoryRepo.GetAllCategories(ctx, page, limit, search, sortBy, sortOrder)
}

func (c *categoryUsecase) GetCategoryByID(ctx context.Context, id uint) (*model.Category, error) {
	return c.categoryRepo.GetCategoryByID(ctx, id)
}

func (c *categoryUsecase) UpdateCategory(ctx context.Context, category *model.Category) error {
	return c.categoryRepo.UpdateCategory(ctx, category)
}

func (c *categoryUsecase) DeleteCategory(ctx context.Context, id uint) error {
	return c.categoryRepo.DeleteCategory(ctx, id)
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepositoryInterface) CategoryUsecaseInterface {
	return &categoryUsecase{categoryRepo: categoryRepo}
}