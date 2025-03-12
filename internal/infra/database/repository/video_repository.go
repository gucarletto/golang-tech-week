package repository

import (
	"context"
	"database/sql"
	"gucarletto/golang-tech-week/internal/domain/entity"
	"gucarletto/golang-tech-week/internal/domain/repository"
	"time"
)

type VideoRepositoryImpl struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) repository.VideoRepository {
	return &VideoRepositoryImpl{
		db: db,
	}
}

func (r *VideoRepositoryImpl) Create(ctx context.Context, video *entity.Video) error {
	query := `INSERT INTO videos (id, title, status, hls_path, upload_status, s3_url, s3_manifest_path, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query, video.ID, video.Title, video.Status, video.HLSPath, video.UploadStatus, video.S3URL, video.S3ManifestPath, video.CreatedAt, video.UpdatedAt)
	return err
}

func (r *VideoRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Video, error) {
	query := `SELECT id, title, status, hls_path, upload_status, s3_url, s3_manifest_path, created_at, updated_at FROM videos WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	video := &entity.Video{}
	err := row.Scan(&video.ID, &video.Title, &video.Status, &video.HLSPath, &video.UploadStatus, &video.S3URL, &video.S3ManifestPath, &video.CreatedAt, &video.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, repository.ErrVideoNotFound
	}
	return video, err
}

func (r *VideoRepositoryImpl) List(ctx context.Context, pageSize int) ([]*entity.Video, error) {
	query := `SELECT id, title, status, hls_path, upload_status, s3_url, s3_manifest_path, created_at, updated_at FROM videos LIMIT $1`
	rows, err := r.db.QueryContext(ctx, query, pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	videos := []*entity.Video{}
	for rows.Next() {
		video := &entity.Video{}
		err := rows.Scan(&video.ID, &video.Title, &video.Status, &video.HLSPath, &video.UploadStatus, &video.S3URL, &video.S3ManifestPath, &video.CreatedAt, &video.UpdatedAt)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func (r *VideoRepositoryImpl) UpdateStatus(ctx context.Context, id, status string) error {
	query := `UPDATE videos SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, status, time.Now(), id)
	return err
}

func (r *VideoRepositoryImpl) UpdateHLSPath(ctx context.Context, id, hlsPath string) error {
	query := `UPDATE videos SET hls_path = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, hlsPath, time.Now(), id)
	return err
}

func (r *VideoRepositoryImpl) UpdateS3Status(ctx context.Context, id, s3Status string) error {
	query := `UPDATE videos SET upload_status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, s3Status, time.Now(), id)
	return err
}

func (r *VideoRepositoryImpl) UpdateS3URLs(ctx context.Context, id, s3URL, s3ManifestURL string) error {
	query := `UPDATE videos SET s3_url = $1, s3_manifest_path = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, s3URL, s3ManifestURL, time.Now(), id)
	return err
}

func (r *VideoRepositoryImpl) UpdateS3Keys(ctx context.Context, id, s3Key, s3ManifestKey string) error {
	query := `UPDATE videos SET s3_url = $1, s3_manifest_path = $2, updated_at = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, s3Key, s3ManifestKey, time.Now(), id)
	return err
}

func (r *VideoRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM videos WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
