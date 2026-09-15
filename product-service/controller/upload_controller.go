package controller

import (
	"pos/product-service/controller/response"
	"pos/product-service/pkg/storage"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type UploadControllerInterface interface {
	UploadProductImage(ctx fiber.Ctx) error
	UploadCategoryImage(ctx fiber.Ctx) error
}

type UploadController struct {
	fileUploadHelper *storage.FileUploadHelper
}

func (u *UploadController) UploadProductImage(ctx fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Errorf("failed to get file: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to upgetload file",
			"error": err.Error(),
		})
	}

	result, err := u.fileUploadHelper.UploadPhoto(ctx.Context(), file, "products")
	if err != nil {
		log.Errorf("failed to upload file: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
			"error": err.Error(),
		})
	}

	response := response.UploadResponse {
		URL: result.URL,
		Path: result.Path,
		Filename: result.Filename,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data": response,
	})
}

func (u *UploadController) UploadCategoryImage(ctx fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Errorf("failed to get file: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to upgetload file",
			"error": err.Error(),
		})
	}

	result, err := u.fileUploadHelper.UploadPhoto(ctx.Context(), file, "categories")
	if err != nil {
		log.Errorf("failed to upload file: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload file",
			"error": err.Error(),
		})
	}

	response := response.UploadResponse {
		URL: result.URL,
		Path: result.Path,
		Filename: result.Filename,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data": response,
	})
}

func NewUploadController(fileUploadHelper *storage.FileUploadHelper) UploadControllerInterface {
	return &UploadController{fileUploadHelper: fileUploadHelper}
}