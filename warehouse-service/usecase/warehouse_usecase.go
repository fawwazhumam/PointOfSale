package usecase

import (
	"context"
	"pos/warehouse-service/model"
	"pos/warehouse-service/repository"
)

type WarehouseUsecaseInterface interface {
	CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) error
	GetAllWarehouses(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.Warehouse, int64, error)
	GetWarehouseByID(ctx context.Context, id uint) (*model.Warehouse, error)
	UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error
	DeleteWarehouse(ctx context.Context, id uint) error
}

type warehouseUsecase struct {
	warehouseRepo repository.WarehouseRepositoryInterface
}

func (w *warehouseUsecase) CreateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	return w.warehouseRepo.CreateWarehouse(ctx, warehouse)
}

func (w *warehouseUsecase) GetAllWarehouses(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.Warehouse, int64, error) {
	return w.warehouseRepo.GetAllWarehouses(ctx, page, limit, search, sortBy, sortOrder)
}

func (w *warehouseUsecase) GetWarehouseByID(ctx context.Context, id uint) (*model.Warehouse, error) {
	return w.warehouseRepo.GetWarehouseByID(ctx, id)
}

func (w *warehouseUsecase) UpdateWarehouse(ctx context.Context, warehouse *model.Warehouse) error {
	return w.warehouseRepo.UpdateWarehouse(ctx, warehouse)
}

func (w *warehouseUsecase) DeleteWarehouse(ctx context.Context, id uint) error {
	return w.warehouseRepo.DeleteWarehouse(ctx, id)
}

func NewWarehouseUsecase(warehouseRepo repository.WarehouseRepositoryInterface) WarehouseUsecaseInterface {
	return &warehouseUsecase{warehouseRepo: warehouseRepo}
}
