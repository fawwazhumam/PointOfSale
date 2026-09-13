package controller

import (
	"pos/product-service/controller/request"
	"pos/product-service/controller/response"
	"pos/product-service/model"
	"pos/product-service/pkg/conv"
	"pos/product-service/pkg/pagination"
	"pos/product-service/pkg/validator"
	"pos/product-service/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type CategoryControllerInterface interface {
	CreateCategory(ctx fiber.Ctx) error
	GetAllCategories(ctx fiber.Ctx) error
	GetCategoryByID(ctx fiber.Ctx) error
	UpdateCategory(ctx fiber.Ctx) error
	DeleteCategory(ctx fiber.Ctx) error
}

type CategoryController struct {
	categoryUsecase usecase.CategoryUsecaseInterface
}

func (c *CategoryController) CreateCategory(ctx fiber.Ctx) error {
	var req request.CreateCategoryRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[CategoryController] CreateCategory - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[CategoryController] CreateCategory - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Category {
		Name: req.Name,
		Tagline: req.Tagline,
		Photo: req.Photo,
	}

	if err := c.categoryUsecase.CreateCategory(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[CategoryController] CreateCategory - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create category",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Category created successfully",
	})

}

func (c *CategoryController) GetAllCategories(ctx fiber.Ctx) error {
	var req request.GetAllCategoryRequest
	if err := ctx.Bind().Query(&req); err != nil {
		log.Errorf("[CategoryController] GetAllCategories - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query parameters",
		})
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	categories, total, err := c.categoryUsecase.GetAllCategories(ctx.Context(), req.Page, req.Limit, req.Search, req.SortBy, req.SortOrder)
	if err != nil {
		log.Errorf("[CategoryController] GetAllCategories - 2: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all categories",
		})
	}

	pagination := pagination.CalculatePagination(req.Page, req.Limit, int(total))

	var categoriesResponse []response.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, response.CategoryResponse{
			ID: category.ID,
			Name: category.Name,
			Tagline: category.Tagline,
			Photo: category.Photo,
		})
	}

	response := response.GetAllCategoriesResponse {
		Categories: categoriesResponse,
		Pagination: pagination,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Categories fetched successfully",
		"data": response,
	})
}

func (c *CategoryController) GetCategoryByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	category, err := c.categoryUsecase.GetCategoryByID(ctx.Context(), idUint)
	if err != nil {
		log.Errorf("[CategoryController] GetCategoryByID - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get category by id",
		})
	}

	response := response.CategoryResponse {
		ID: category.ID,
		Name: category.Name,
		Tagline: category.Tagline,
		Photo: category.Photo,
		CountProduct: len(category.Products),
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category fetched successfulluy",
		"data": response,
	})
}

func (c *CategoryController) UpdateCategory(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	var req request.CreateCategoryRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[CategoryController] UpdateCategory - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[CategoryController] UpdateCategory - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Category {
		ID: idUint,
		Name: req.Name,
		Tagline: req.Tagline,
		Photo: req.Photo,
	}

	if err := c.categoryUsecase.UpdateCategory(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[CategoryController] UpdateCategory - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to UpdateCategory",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category updated successfully",
	})
}

func (c *CategoryController) DeleteCategory(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	if err := c.categoryUsecase.DeleteCategory(ctx.Context(), idUint); err != nil {
		log.Errorf("[CategoryController] DeleteCategory - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delelete category",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category delete successfully",
	})
}

func NewCategoryController(categoryUsecase usecase.CategoryUsecaseInterface) CategoryControllerInterface {
	return &CategoryController{categoryUsecase: categoryUsecase}
}