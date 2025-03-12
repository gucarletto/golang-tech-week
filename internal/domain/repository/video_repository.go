package repository

import (
	"context"
	"errors"
	"gucarletto/golang-tech-week/internal/domain/entity"
)

var (
	ErrVideoNotFound = errors.New("video not found")
)

type VideoRepository interface {
	Create(ctx context.Context, video *entity.Video) error
	FindByID(ctx context.Context, id string) (*entity.Video, error)
	List(ctx context.Context, pageSize int) ([]*entity.Video, error)
	UpdateStatus(ctx context.Context, id, status string) error
	UpdateHLSPath(ctx context.Context, id, hlsPath string) error
	UpdateS3Status(ctx context.Context, id, s3Status string) error
	UpdateS3URLs(ctx context.Context, id, s3URL, s3ManifestURL string) error
	UpdateS3Keys(ctx context.Context, id, s3Key, s3ManifestKey string) error
	Delete(ctx context.Context, id string) error
}
