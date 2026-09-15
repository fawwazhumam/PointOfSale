package repository

import (
	"context"
	"errors"
	"pos/warehouse-service/model"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

type WarehouseProductRepositoryInterface interface {
	CreateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error
	GetAllWarehouseProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.WarehouseProduct, int64, error)
	GetWarehouseProductByID(ctx context.Context, id uint) (*model.WarehouseProduct, error)
	UpdateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error
	DeleteWarehouseProduct(ctx context.Context, id uint) error
}

type warehouseProductRepository struct {
	db *gorm.DB
}

func (w *warehouseProductRepository) CreateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error {
	select {
	case <-ctx.Done():
		log.Errorf("[WarehouseProductRepository] CreateWarehouseProduct - 1: %v", ctx.Err())
		return ctx.Err()
	default:
		var count int64
		if err := w.db.WithContext(ctx).Model(&model.WarehouseProduct{}).
			Where("warehouse_id = ? AND product_id = ?", warehouseProduct.WarehouseID, warehouseProduct.ProductID).
			Count(&count).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] CreateWarehouseProduct - 2: %v", err)
			return err
		}
		if count > 0 {
			return errors.New("this product is already assigned to this warehouse")
		}

		return w.db.WithContext(ctx).Create(warehouseProduct).Error
	}
}

func (w *warehouseProductRepository) GetAllWarehouseProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.WarehouseProduct, int64, error) {
	select {
	case <-ctx.Done():
		log.Errorf("[WarehouseProductRepository] GetAllWarehouseProducts - 1: %v", ctx.Err())
		return nil, 0, ctx.Err()
	default:
		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 10
		}
		if sortBy == "" {
			sortBy = "created_at"
		}
		if sortOrder == "" {
			sortOrder = "desc"
		}

		// calculate offset
		offset := (page - 1) * limit

		query := w.db.WithContext(ctx).Model(&model.WarehouseProduct{})

		if search != "" {
			query = query.Joins("JOIN warehouses ON warehouses.id = warehouse_products.warehouse_id").
				Where("warehouses.name ILIKE ?", "%"+search+"%")
		}
		// ILIKE meniadakan case sensitive

		var totalRecords int64
		if err := query.Count(&totalRecords).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] GetAllWarehouseProducts - 2: Failed to count warehouse products %v", err)
			return nil, 0, err
		}

		var warehouseProducts []model.WarehouseProduct
		if err := query.Order(sortBy + " " + sortOrder).Preload("Warehouse").Offset(offset).Limit(limit).Find(&warehouseProducts).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] GetAllWarehouseProducts - 3: Failed to get all warehouse products: %v", err)
			return nil, 0, err
		}

		return warehouseProducts, totalRecords, nil
	}
}

func (w *warehouseProductRepository) GetWarehouseProductByID(ctx context.Context, id uint) (*model.WarehouseProduct, error) {
	select {
	case <-ctx.Done():
		log.Errorf("[WarehouseProductRepository] GetWarehouseProductByID - 1: %v", ctx.Err())
		return nil, ctx.Err()
	default:
		warehouseProduct := model.WarehouseProduct{}
		if err := w.db.WithContext(ctx).Preload("Warehouse").Where("id = ?", id).First(&warehouseProduct).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] GetWarehouseProductByID - 2: %v", err)
			return nil, err
		}
		return &warehouseProduct, nil
	}
}

func (w *warehouseProductRepository) UpdateWarehouseProduct(ctx context.Context, warehouseProduct *model.WarehouseProduct) error {
	select {
	case <-ctx.Done():
		log.Errorf("[WarehouseProductRepository] UpdateWarehouseProduct - 1: %v", ctx.Err())
		return ctx.Err()
	default:
		existingWarehouseProduct := model.WarehouseProduct{}
		if err := w.db.WithContext(ctx).Where("id = ?", warehouseProduct.ID).First(&existingWarehouseProduct).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] UpdateWarehouseProduct - 2: %v", err)
			return err
		}

		existingWarehouseProduct.WarehouseID = warehouseProduct.WarehouseID
		existingWarehouseProduct.ProductID = warehouseProduct.ProductID
		existingWarehouseProduct.Stock = warehouseProduct.Stock

		return w.db.WithContext(ctx).Save(&existingWarehouseProduct).Error
	}
}

func (w *warehouseProductRepository) DeleteWarehouseProduct(ctx context.Context, id uint) error {
	select {
	case <-ctx.Done():
		log.Errorf("[WarehouseProductRepository] DeleteWarehouseProduct - 1: %v", ctx.Err())
		return ctx.Err()
	default:
		warehouseProduct := model.WarehouseProduct{}
		if err := w.db.WithContext(ctx).Where("id = ?", id).First(&warehouseProduct).Error; err != nil {
			log.Errorf("[WarehouseProductRepository] DeleteWarehouseProduct - 2: %v", err)
			return err
		}
		return w.db.WithContext(ctx).Delete(&warehouseProduct).Error
	}
}

func NewWarehouseProductRepository(db *gorm.DB) WarehouseProductRepositoryInterface {
	return &warehouseProductRepository{db: db}
}
