package app

import (
	"log"
	"pos/product-service/configs"
	"pos/product-service/controller"
	"pos/product-service/database"
	"pos/product-service/pkg/storage"
	"pos/product-service/repository"
	"pos/product-service/usecase"
)

type Container struct {
	ProductController controller.ProductControllerInterface
	CategoryController controller.CategoryControllerInterface
	UploadController controller.UploadControllerInterface
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	categoryRepo := repository.NewCategoryRepository(db.DB)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryController := controller.NewCategoryController(categoryUsecase)

	productRepo := repository.NewProductRepository(db.DB)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productController := controller.NewProductController(productUsecase)

	supabaseStorage := storage.NewSupabaseStorage(*config)
	fileUploadHelper := storage.NewFileUploadHelper(supabaseStorage, *config)
	uploadController := controller.NewUploadController(fileUploadHelper)

	return &Container{
		ProductController: productController,
		CategoryController: categoryController,
		UploadController: uploadController,
	}
}