package services

import (
	"context"
	"pos-app/backend/internal/models"

	"github.com/google/uuid"
)

// FileStorageService menangani upload/download file dan manajemen file di storage (lokal/S3/dll).
type FileStorageService interface {
	UploadFile(ctx context.Context, fileData []byte, fileName string, uploaderID uuid.UUID) (*models.FileMetadata, error)
	DownloadFile(ctx context.Context, fileID uuid.UUID) ([]byte, *models.FileMetadata, error)
	DeleteFile(ctx context.Context, fileID uuid.UUID) error
	ListFiles(ctx context.Context, ownerID uuid.UUID) ([]*models.FileMetadata, error)
	GetFileMetadata(ctx context.Context, fileID uuid.UUID) (*models.FileMetadata, error)
	ShareFileLink(ctx context.Context, fileID uuid.UUID) (string, error)
	ArchiveFile(ctx context.Context, fileID uuid.UUID) error
	RestoreFile(ctx context.Context, fileID uuid.UUID) error
}
