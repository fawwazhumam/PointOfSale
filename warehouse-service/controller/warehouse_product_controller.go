package controller

import (
	"pos/warehouse-service/controller/request"
	"pos/warehouse-service/controller/response"
	"pos/warehouse-service/model"
	"pos/warehouse-service/pkg/conv"
	"pos/warehouse-service/pkg/pagination"
	"pos/warehouse-service/pkg/validator"
	"pos/warehouse-service/usecase"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type WarehouseProductControllerInterface interface {
	CreateWarehouseProduct(ctx fiber.Ctx) error
	GetAllWarehouseProducts(ctx fiber.Ctx) error
	GetWarehouseProductByID(ctx fiber.Ctx) error
	UpdateWarehouseProduct(ctx fiber.Ctx) error
	DeleteWarehouseProduct(ctx fiber.Ctx) error
}

type WarehouseProductController struct {
	warehouseProductUsecase usecase.WarehouseProductUsecaseInterface
}

func (w *WarehouseProductController) CreateWarehouseProduct(ctx fiber.Ctx) error {
	var req request.CreateWarehouseProductRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.WarehouseProduct{
		WarehouseID: req.WarehouseID,
		ProductID:   req.ProductID,
		Stock:       req.Stock,
	}

	if err := w.warehouseProductUsecase.CreateWarehouseProduct(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[WarehouseProductController] CreateWarehouseProduct - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create warehouse product",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Warehouse product created successfully",
	})
}

func (w *WarehouseProductController) GetAllWarehouseProducts(ctx fiber.Ctx) error {
	var req request.GetAllWarehouseProductRequest
	if err := ctx.Bind().Query(&req); err != nil {
		log.Errorf("[WarehouseProductController] GetAllWarehouseProducts - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid query parameters",
		})
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	warehouseProducts, total, err := w.warehouseProductUsecase.GetAllWarehouseProducts(ctx.Context(), req.Page, req.Limit, req.Search, req.SortBy, req.SortOrder)
	if err != nil {
		log.Errorf("[WarehouseProductController] GetAllWarehouseProducts - 2: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all warehouse products",
		})
	}

	paginationInfo := pagination.CalculatePagination(req.Page, req.Limit, int(total))

	var warehouseProductResponse []response.WarehouseProductResponse
	for _, warehouseProduct := range warehouseProducts {
		warehouseProductResponse = append(warehouseProductResponse, response.WarehouseProductResponse{
			ID:            warehouseProduct.ID,
			WarehouseID:   warehouseProduct.WarehouseID,
			WarehouseName: warehouseProduct.Warehouse.Name,
			ProductID:     warehouseProduct.ProductID,
			Stock:         warehouseProduct.Stock,
		})
	}

	result := response.GetAllWarehouseProductsResponse{
		WarehouseProducts: warehouseProductResponse,
		Pagination:        paginationInfo,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse products fetched successfully",
		"data":    result,
	})
}

func (w *WarehouseProductController) GetWarehouseProductByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	warehouseProduct, err := w.warehouseProductUsecase.GetWarehouseProductByID(ctx.Context(), idUint)
	if err != nil {
		log.Errorf("[WarehouseProductController] GetWarehouseProductByID - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get warehouse product by id",
		})
	}

	result := response.WarehouseProductResponse{
		ID:            warehouseProduct.ID,
		WarehouseID:   warehouseProduct.WarehouseID,
		WarehouseName: warehouseProduct.Warehouse.Name,
		ProductID:     warehouseProduct.ProductID,
		Stock:         warehouseProduct.Stock,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse product fetched successfully",
		"data":    result,
	})
}

func (w *WarehouseProductController) UpdateWarehouseProduct(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	var req request.CreateWarehouseProductRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[WarehouseProductController] UpdateWarehouseProduct - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[WarehouseProductController] UpdateWarehouseProduct - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.WarehouseProduct{
		ID:          idUint,
		WarehouseID: req.WarehouseID,
		ProductID:   req.ProductID,
		Stock:       req.Stock,
	}

	if err := w.warehouseProductUsecase.UpdateWarehouseProduct(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[WarehouseProductController] UpdateWarehouseProduct - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update warehouse product",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse product updated successfully",
	})
}

func (w *WarehouseProductController) DeleteWarehouseProduct(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	if err := w.warehouseProductUsecase.DeleteWarehouseProduct(ctx.Context(), idUint); err != nil {
		log.Errorf("[WarehouseProductController] DeleteWarehouseProduct - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete warehouse product",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse product deleted successfully",
	})
}

func NewWarehouseProductController(warehouseProductUsecase usecase.WarehouseProductUsecaseInterface) WarehouseProductControllerInterface {
	return &WarehouseProductController{warehouseProductUsecase: warehouseProductUsecase}
}
