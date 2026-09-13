package usecase

import (
	"context"
	"pos/product-service/model"
	"pos/product-service/repository"
)

type ProductUsecaseInterface interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	GetAllProducts(ctx context.Context, page, limit int, search, sortBy, sortOrder string) ([]model.Product, int64, error)
	GetProductByID(ctx context.Context, id uint) (*model.Product, error)
	GetProductByBarcode(ctx context.Context, barcode string) (*model.Product, error)
	UpdateProduct(ctx context.Context, product *model.Product) error
	DeleteProduct(ctx context.Context, id uint) error
}

type productUsecase struct {
	productRepo repository.ProductRepositoryInterface
}

func (p *productUsecase) CreateProduct(ctx context.Context, product *model.Product) error {
	return p.productRepo.CreateProduct(ctx, product)
}

func (p *productUsecase) GetAllProducts(ctx context.Context, page int, limit int, search string, sortBy string, sortOrder string) ([]model.Product, int64, error) {
	return p.productRepo.GetAllProducts(ctx, page, limit, search, sortBy, sortOrder)
}

func (p *productUsecase) GetProductByID(ctx context.Context, id uint) (*model.Product, error) {
	return p.productRepo.GetProductByID(ctx, id)
}

func (p *productUsecase) GetProductByBarcode(ctx context.Context, barcode string) (*model.Product, error) {
	return p.productRepo.GetProductByBarcode(ctx, barcode)
}

func (p *productUsecase) UpdateProduct(ctx context.Context, product *model.Product) error {
	return p.productRepo.UpdateProduct(ctx, product)
}

func (p *productUsecase) DeleteProduct(ctx context.Context, id uint) error {
	return p.productRepo.DeleteProduct(ctx, id)

	// harus ada djustment untuk menghapus product oada merchant dan warehouse juga karena ada assign product disana
}

func NewProductUsecase(productRepo repository.ProductRepositoryInterface) ProductUsecaseInterface {
	return &productUsecase{productRepo: productRepo}
}