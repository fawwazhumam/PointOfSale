package usecase

import (
	"context"
	"pos/warehouse-service/model"
	"pos/warehouse-service/repository"
)

type WarehouseProductUsecaseInterface interface {
	CreateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error
	GetAllWarehouseProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.WarehouseProduct, int64, error)
	GetWarehouseProductByID(ctx context.Context, id uint) (*model.WarehouseProduct, error)
	UpdateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error
	DeleteWarehouseProduct(ctx context.Context, id uint) error
}

type warehouseProductUsecase struct {
	warehouseProductRepo repository.WarehouseProductRepositoryInterface
}

func (w *warehouseProductUsecase) CreateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error {
	return w.warehouseProductRepo.CreateWarehouseProduct(ctx, warehouseProduct)
}

func (w *warehouseProductUsecase) GetAllWarehouseProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.WarehouseProduct, int64, error) {
	return w.warehouseProductRepo.GetAllWarehouseProducts(ctx, page, limit, search, sortBy, sortOrder)
}

func (w *warehouseProductUsecase) GetWarehouseProductByID(ctx context.Context, id uint) (*model.WarehouseProduct, error) {
	return w.warehouseProductRepo.GetWarehouseProductByID(ctx, id)
}

func (w *warehouseProductUsecase) UpdateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error {
	return w.warehouseProductRepo.UpdateWarehouseProduct(ctx, warehouseProduct)
}

func (w *warehouseProductUsecase) DeleteWarehouseProduct(ctx context.Context, id uint) error {
	return w.warehouseProductRepo.DeleteWarehouseProduct(ctx, id)
}

func NewWarehouseProductUsecase(warehouseProductRepo repository.WarehouseProductRepositoryInterface) WarehouseProductUsecaseInterface {
	return &warehouseProductUsecase{warehouseProductRepo: warehouseProductRepo}
}
