package entity

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// Videos statuses
const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusConverting = "converting"
	StatusUploading  = "uploading"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
)

const (
	UploadStatusNonde       = "none"
	UploadStatusPendingS3   = "pending_s3"
	UploadStatusUploadingS3 = "uploading_s3"
	UploadStatusCompletedS3 = "completed_s3"
	UploadStatusFailedS3    = "failed_s3"
)

const (
	FileTypeManifest = "manifest"
	FileTypeSegment  = "segment"
)

// Video represents a video entity
type Video struct {
	ID             string
	Title          string
	FilePath       string
	HLSPath        string
	ManifestPath   string
	S3ManifestPath string
	S3URL          string
	Status         string
	UploadStatus   string
	ErrorMessage   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewVideo creates a new video entity
func NewVideo(title, filePath string) *Video {
	return &Video{
		ID:        uuid.New().String(),
		Title:     title,
		FilePath:  filePath,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// MarksAsProcessing sets the video status to processing
func (v *Video) MarksAsProcessing() {
	v.Status = StatusProcessing
	v.UpdatedAt = time.Now()
}

// MarkAsCompleted sets the video status to completed
func (v *Video) MarkAsCompleted(hlsPath, manifestPath string) {
	v.Status = StatusCompleted
	v.UpdatedAt = time.Now()
	v.HLSPath = hlsPath
	v.ManifestPath = manifestPath
}

// MarkAsFailed sets the video status to failed and records an error message
func (v *Video) MarkAsFailed(errorMessage string) {
	v.Status = StatusFailed
	v.ErrorMessage = errorMessage
	v.UpdatedAt = time.Now()
}

// SetS3URL sets the S3 URL for the video
func (v *Video) SetS3URL(url string) {
	v.S3URL = url
	v.UpdatedAt = time.Now()
}

// SetS3ManifestURL sets the S3 manifest URL for the video
func (v *Video) SetS3ManifestURL(url string) {
	v.S3ManifestPath = url
	v.UpdatedAt = time.Now()
}

// IsCompleted checks if the video status is completed
func (v *Video) IsCompleted() bool {
	return v.Status == StatusCompleted
}

// GetHLSDirectory returns the directory of the HLS files
func (v *Video) GetHLSDirectory() string {
	return filepath.Dir(v.HLSPath)
}

// GetManifestPath returns the manifest path of the video
func (v *Video) GetManifestPath() string {
	return v.ManifestPath
}

// GenerateOutputPath generates the output path for the video based on the base directory
func (v *Video) GenerateOutputPath(baseDir string) string {
	return filepath.Join(baseDir, "/converted/", v.ID)
}
