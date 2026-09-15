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

type WarehouseControllerInterface interface {
	CreateWarehouse(ctx fiber.Ctx) error
	GetAllWarehouses(ctx fiber.Ctx) error
	GetWarehouseByID(ctx fiber.Ctx) error
	UpdateWarehouse(ctx fiber.Ctx) error
	DeleteWarehouse(ctx fiber.Ctx) error
}

type WarehouseController struct {
	warehouseUsecase usecase.WarehouseUsecaseInterface
}

func (w *WarehouseController) CreateWarehouse(ctx fiber.Ctx) error {
	var req request.CreateWarehouseRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[WarehouseController] CreateWarehouse - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[WarehouseController] CreateWarehouse - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Warehouse{
		Name:    req.Name,
		Address: req.Address,
		Photo:   req.Photo,
		Phone:   req.Phone,
	}

	if err := w.warehouseUsecase.CreateWarehouse(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[WarehouseController] CreateWarehouse - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create warehouse",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Warehouse created successfully",
	})
}

func (w *WarehouseController) GetAllWarehouses(ctx fiber.Ctx) error {
	var req request.GetAllWarehouseRequest
	if err := ctx.Bind().Query(&req); err != nil {
		log.Errorf("[WarehouseController] GetAllWarehouses - 1: %v", err)
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

	warehouses, total, err := w.warehouseUsecase.GetAllWarehouses(ctx.Context(), req.Page, req.Limit, req.Search, req.SortBy, req.SortOrder)
	if err != nil {
		log.Errorf("[WarehouseController] GetAllWarehouses - 2: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get all warehouses",
		})
	}

	paginationInfo := pagination.CalculatePagination(req.Page, req.Limit, int(total))

	var warehouseResponse []response.WarehouseResponse
	for _, warehouse := range warehouses {
		warehouseResponse = append(warehouseResponse, response.WarehouseResponse{
			ID:        warehouse.ID,
			Name:      warehouse.Name,
			Address:   warehouse.Address,
			Photo:     warehouse.Photo,
			Phone:     warehouse.Phone,
			CreatedAt: warehouse.CreatedAt,
		})
	}

	result := response.GetAllWarehousesResponse{
		Warehouses: warehouseResponse,
		Pagination: paginationInfo,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouses fetched successfully",
		"data":    result,
	})
}

func (w *WarehouseController) GetWarehouseByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	warehouse, err := w.warehouseUsecase.GetWarehouseByID(ctx.Context(), idUint)
	if err != nil {
		log.Errorf("[WarehouseController] GetWarehouseByID - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get warehouse by id",
		})
	}

	result := response.WarehouseResponse{
		ID:        warehouse.ID,
		Name:      warehouse.Name,
		Address:   warehouse.Address,
		Photo:     warehouse.Photo,
		Phone:     warehouse.Phone,
		CreatedAt: warehouse.CreatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse fetched successfully",
		"data":    result,
	})
}

func (w *WarehouseController) UpdateWarehouse(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	var req request.CreateWarehouseRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[WarehouseController] UpdateWarehouse - 1: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(req); err != nil {
		log.Errorf("[WarehouseController] UpdateWarehouse - 2: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	reqModel := model.Warehouse{
		ID:      idUint,
		Name:    req.Name,
		Address: req.Address,
		Photo:   req.Photo,
		Phone:   req.Phone,
	}

	if err := w.warehouseUsecase.UpdateWarehouse(ctx.Context(), &reqModel); err != nil {
		log.Errorf("[WarehouseController] UpdateWarehouse - 3: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update warehouse",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse updated successfully",
	})
}

func (w *WarehouseController) DeleteWarehouse(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	idUint := conv.StringToUint(id)

	if err := w.warehouseUsecase.DeleteWarehouse(ctx.Context(), idUint); err != nil {
		log.Errorf("[WarehouseController] DeleteWarehouse - 1: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete warehouse",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Warehouse deleted successfully",
	})
}

func NewWarehouseController(warehouseUsecase usecase.WarehouseUsecaseInterface) WarehouseControllerInterface {
	return &WarehouseController{warehouseUsecase: warehouseUsecase}
}
