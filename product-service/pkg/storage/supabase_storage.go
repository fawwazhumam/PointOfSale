package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"pos/product-service/configs"

	"strings"
	"time"

	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseInterface interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error)
}

type SupabaseStorage struct {
	client *storage_go.Client
	cfg    configs.Config
}

// upload file
func (s *SupabaseStorage) UploadFile(ctx context.Context, file *multipart.FileHeader, folder string) (*UploadResult, error) {
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	// generate unique filename
	ext := filepath.Ext(file.Filename)
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%s_%d%s", strings.TrimSuffix(file.Filename, ext), timestamp, ext)

	filePath := fmt.Sprintf("%s/%s", folder, filename)

	// prioritaskan deteksi dari ekstensi file, karena header Content-Type dari
	// client sering generic/tidak akurat (misal "application/octet-stream")
	var contentType string
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	case ".svg":
		contentType = "image/svg+xml"
	}

	if contentType == "" {
		contentType = file.Header.Get("Content-Type")
	}
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// upload file using the client set up in NewSupabaseStorage
	_, err = s.client.UploadFile(s.cfg.Supabase.Bucket, filePath, src, storage_go.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to supabase: %w", err)
	}

	// Get public URL
	publicUrl := s.client.GetPublicUrl(s.cfg.Supabase.Bucket, filePath)

	return &UploadResult{
		URL:      publicUrl.SignedURL,
		Path:     filePath,
		Filename: filename,
	}, nil
}

type UploadResult struct {
	URL      string `json:"url"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
}



func NewSupabaseStorage(cfg configs.Config) SupabaseInterface {
	client := storage_go.NewClient(cfg.Supabase.Url, cfg.Supabase.Key, map[string]string{
		"apikey": cfg.Supabase.Key,
	})
	return &SupabaseStorage{
		client: client,
		cfg:    cfg,
	}
}
