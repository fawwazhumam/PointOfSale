package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"pos/user-service/configs"
	"strings"
)

const (
	MaxImageSize = 2 * 1024 * 1024 // 2mb

	AllowedImageExtentions = ".jpg,.jpeg,.png,.webp,.svg"
)

type FileUploadHelper struct {
	storage SupabaseInterface
	cfg     configs.Config
}

func NewFileUploadHelper(storage SupabaseInterface, cfg configs.Config) *FileUploadHelper {
	return &FileUploadHelper{
		storage: storage,
		cfg: cfg,
	}
}

func (h *FileUploadHelper) UploadPhoto(ctx context.Context, file *multipart.FileHeader) (*UploadResult, error) {
	if err := h.validateImageFile(file, MaxImageSize); err != nil {
		log.Fatalf("failed to validate image file: %v", err)
		return nil, err
	}

	result, err := h.storage.UploadFile(ctx, file, "users")
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
		return nil, err
	}

	return result, nil
}

func (h *FileUploadHelper) validateImageFile(file *multipart.FileHeader, maxSize int64) error {
	if !validateFileSize(file.Size, maxSize) {
		return fmt.Errorf("file size exceeds the maximum allowed size")
	}

	if !validateFileExtention(getFileExtention(file.Filename), AllowedImageExtentions) {
		return fmt.Errorf("file typenya kocak")
	}

	return nil
}

func validateFileSize(size int64, maxSize int64) bool {
	return size <= maxSize
}

func getFileExtention(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}

func validateFileExtention(extention string, allowedExtentions string) bool {
	allowed := strings.Split(allowedExtentions, ",")
	for _, ext := range allowed {
		if strings.TrimSpace(ext) == extention {
			return true
		}
	}

	return false
}
