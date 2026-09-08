package app

import (
	"log"
	"pos/user-service/configs"
	"pos/user-service/controller"
	"pos/user-service/database"
	"pos/user-service/pkg/storage"
	"pos/user-service/repository"
	"pos/user-service/service"
	"pos/user-service/usecase"
)

type Container struct {
	RoleController controller.RoleControllerInterface
	UserController controller.UserControllerInterface
	AuthController controller.AuthControllerInterface
	UploadController controller.UploadControllerInterface
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.ConnectionPostgres(*config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	supabaseStorage := storage.NewSupabaseStorage(*config)

	fileUploadHelper := storage.NewFileUploadHelper(supabaseStorage, *config)

	roleRepo := repository.NewRoleRepository(db.DB)
	roleUsecase := usecase.NewRoleUseCase(roleRepo)
	roleController := controller.NewRoleController(roleUsecase)

	userRepo := repository.NewUserRepository(db.DB)
	rabbitMQService, err := service.NewRabbitMQService(*config)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	userUsecase := usecase.NewUserUsecase(userRepo, rabbitMQService)
	userController := controller.NewUserController(userUsecase)

	authController := controller.NewAuthController(userUsecase)
	uploadController := controller.NewUploadController(fileUploadHelper)

	return &Container{
		RoleController: roleController,
		UserController: userController,
		AuthController: authController,
		UploadController: uploadController,
	}
}