package controller

import (
	"pos/product-service/controller/request"
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
		log.Errorf("[ProductController]")
	}
}