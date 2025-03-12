package repository

import (
	"context"
	"testing"
	"time"

	"gucarletto/golang-tech-week/internal/domain/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	video := &entity.Video{
		ID:             "1",
		Title:          "Test Video",
		Status:         "uploaded",
		HLSPath:        "/path/to/hls",
		UploadStatus:   "completed",
		S3URL:          "http://s3.url",
		S3ManifestPath: "http://s3.manifest.url",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	mock.ExpectExec(`INSERT INTO videos`).
		WithArgs(video.ID, video.Title, video.Status, video.HLSPath, video.UploadStatus, video.S3URL, video.S3ManifestPath, video.CreatedAt, video.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(context.Background(), video)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	video := &entity.Video{
		ID:             "1",
		Title:          "Test Video",
		Status:         "uploaded",
		HLSPath:        "/path/to/hls",
		UploadStatus:   "completed",
		S3URL:          "http://s3.url",
		S3ManifestPath: "http://s3.manifest.url",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "status", "hls_path", "upload_status", "s3_url", "s3_manifest_path", "created_at", "updated_at"}).
		AddRow(video.ID, video.Title, video.Status, video.HLSPath, video.UploadStatus, video.S3URL, video.S3ManifestPath, video.CreatedAt, video.UpdatedAt)

	mock.ExpectQuery(`SELECT id, title, status, hls_path, upload_status, s3_url, s3_manifest_path, created_at, updated_at FROM videos WHERE id = \$1`).
		WithArgs(video.ID).
		WillReturnRows(rows)

	result, err := repo.FindByID(context.Background(), video.ID)
	assert.NoError(t, err)
	assert.Equal(t, video, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateStatus(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"
	status := "processed"

	mock.ExpectExec(`UPDATE videos SET status = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(status, sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateStatus(context.Background(), id, status)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateHLSPath(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"
	hlsPath := "/new/path/to/hls"

	mock.ExpectExec(`UPDATE videos SET hls_path = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(hlsPath, sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateHLSPath(context.Background(), id, hlsPath)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateS3Status(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"
	s3Status := "in_progress"

	mock.ExpectExec(`UPDATE videos SET upload_status = \$1, updated_at = \$2 WHERE id = \$3`).
		WithArgs(s3Status, sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateS3Status(context.Background(), id, s3Status)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateS3URLs(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"
	s3URL := "http://new.s3.url"
	s3ManifestURL := "http://new.s3.manifest.url"

	mock.ExpectExec(`UPDATE videos SET s3_url = \$1, s3_manifest_path = \$2, updated_at = \$3 WHERE id = \$4`).
		WithArgs(s3URL, s3ManifestURL, sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateS3URLs(context.Background(), id, s3URL, s3ManifestURL)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateS3Keys(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"
	s3Key := "new_s3_key"
	s3ManifestKey := "new_s3_manifest_key"

	mock.ExpectExec(`UPDATE videos SET s3_url = \$1, s3_manifest_path = \$2, updated_at = \$3 WHERE id = \$4`).
		WithArgs(s3Key, s3ManifestKey, sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateS3Keys(context.Background(), id, s3Key, s3ManifestKey)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewVideoRepository(db)
	id := "1"

	mock.ExpectExec(`DELETE FROM videos WHERE id = \$1`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(context.Background(), id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
