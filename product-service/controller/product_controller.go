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

type ProductControllerInterface interface {
	CreateProduct(ctx fiber.Ctx) error
	GetAllProducts(ctx fiber.Ctx) error
	GetProductByID(ctx fiber.Ctx) error
	GetProductByBarcode(ctx fiber.Ctx) error
	UpdateProduct(ctx fiber.Ctx) error
	DeleteProduct(ctx fiber.Ctx) error
}

type productController struct {
	productUsecase usecase.ProductUsecaseInterface
}

func (p *productController) CreateProduct(ctx fiber.Ctx) error {
	var req request.CreateProductRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[ProductController] CreateProduct - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[ProductController] CreateProduct - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Product {
		Name: req.Name,
		Barcode: req.Barcode,
		CategoryID: req.CategoryID,
		Thumbnail: req.Thumbnail,
		About: req.About,
		Price: float64(req.Price),
		IsPopular: req.IsPopular,
	}

	if err := p.productUsecase.CreateProduct(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[ProductController] CreateProduct - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
	})

}

func (p *productController) GetAllProducts(ctx fiber.Ctx) error {
	var req request.GetAllProductRequest
	if err := ctx.Bind().Query(&req); err != nil {
		log.Errorf("[ProductController] GetAllProduct - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	products, total, err := p.productUsecase.GetAllProducts(ctx.Context(), req.Page, req.Limit, req.Search, req.SortBy, req.SortOrder)
	if err != nil {
		log.Errorf("[ProductController] GetAllProduct - 2: %v", err)
		return  ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all products",
		})
	}

	pagination := pagination.CalculatePagination(req.Page, req.Limit, int(total))
	var productResponse []response.ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, response.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			Barcode: product.Barcode,
			CategoryID: product.CategoryID,
			Thumbnail: product.Thumbnail,
			About: product.About,
			Price: int(product.Price),
			IsPopular: product.IsPopular,
			Category: response.CategoryResponse{
				ID: product.Category.ID,
				Name: product.Category.Name,
				Tagline: product.Category.Tagline,
				Photo: product.Category.Photo,
			},
		})
	}

	response := response.GetAllProductResponse {
		Products: productResponse,
		Pagination: pagination,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product fetched successfully",
		"data": response,
	})
}

func (p *productController) GetProductByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	product, err := p.productUsecase.GetProductByID(ctx.Context(), idUint)
	if err != nil {
		log.Errorf("[ProductController] GetProductByID - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get product by barcode",
		})
	}

	response := response.ProductResponse {
		ID: product.ID,
		Name: product.Name,
		Barcode: product.Barcode,
		CategoryID: product.CategoryID,
		Thumbnail: product.Thumbnail,
		About: product.About,
		Price: int(product.Price),
		IsPopular: product.IsPopular,
		Category: response.CategoryResponse{
			ID: product.Category.ID,
			Name: product.Category.Name,
			Tagline: product.Category.Tagline,
			Photo: product.Category.Photo,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product fetched successfully",
		"data": response,
	})
}

func (p *productController) GetProductByBarcode(ctx fiber.Ctx) error {
	barcode := ctx.Params("barcode")

	product, err := p.productUsecase.GetProductByBarcode(ctx.Context(), barcode)
	if err != nil {
		log.Errorf("[ProductController] GetProductByBarcode - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get product by barcode",
		})
	}

	response := response.ProductResponse {
		ID: product.ID,
		Name: product.Name,
		Barcode: product.Barcode,
		CategoryID: product.CategoryID,
		Thumbnail: product.Thumbnail,
		About: product.About,
		Price: int(product.Price),
		IsPopular: product.IsPopular,
		Category: response.CategoryResponse{
			ID: product.Category.ID,
			Name: product.Category.Name,
			Tagline: product.Category.Tagline,
			Photo: product.Category.Photo,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product fetched successfully",
		"data": response,
	})
}

func (p *productController) UpdateProduct(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	var req request.CreateProductRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[ProductController] UpdateProduct - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[ProductController] UpdateProduct - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Product {
		ID: idUint,
		Name: req.Name,
		Barcode: req.Barcode,
		CategoryID: req.CategoryID,
		Thumbnail: req.Thumbnail,
		About: req.About,
		Price: float64(req.Price),
		IsPopular: req.IsPopular,
	}

	if err := p.productUsecase.UpdateProduct(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[ProductController] UpdateProduct - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
	})
}

func (p *productController) DeleteProduct(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	if err := p.productUsecase.DeleteProduct(ctx.Context(), idUint); err != nil {
		log.Errorf("[ProductController] DeleteProduct - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete product",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}

func NewProductController(productUsecase usecase.ProductUsecaseInterface) ProductControllerInterface {
	return &productController{productUsecase: productUsecase}
}