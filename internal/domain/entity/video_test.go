package entity

import (
	"path/filepath"
	"testing"
)

func TestNewVideo(t *testing.T) {
	title := "Test Video"
	filePath := "/path/to/video.mp4"
	video := NewVideo(title, filePath)

	if video.Title != title {
		t.Errorf("expected title %s, got %s", title, video.Title)
	}
	if video.FilePath != filePath {
		t.Errorf("expected filePath %s, got %s", filePath, video.FilePath)
	}
	if video.Status != StatusPending {
		t.Errorf("expected status %s, got %s", StatusPending, video.Status)
	}
	if video.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}
	if video.UpdatedAt.IsZero() {
		t.Error("expected UpdatedAt to be set")
	}
}

func TestMarksAsProcessing(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	video.MarksAsProcessing()

	if video.Status != StatusProcessing {
		t.Errorf("expected status %s, got %s", StatusProcessing, video.Status)
	}
	if video.UpdatedAt.Before(video.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

func TestMarkAsCompleted(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	video.MarkAsCompleted(video.HLSPath, video.ManifestPath)

	if video.Status != StatusCompleted {
		t.Errorf("expected status %s, got %s", StatusCompleted, video.Status)
	}
	if video.UpdatedAt.Before(video.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

func TestMarkAsFailed(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	errorMessage := "Test error"
	video.MarkAsFailed(errorMessage)

	if video.Status != StatusFailed {
		t.Errorf("expected status %s, got %s", StatusFailed, video.Status)
	}
	if video.ErrorMessage != errorMessage {
		t.Errorf("expected errorMessage %s, got %s", errorMessage, video.ErrorMessage)
	}
	if video.UpdatedAt.Before(video.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

func TestSetS3URL(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	s3URL := "https://s3.amazonaws.com/bucket/video.mp4"
	video.SetS3URL(s3URL)

	if video.S3URL != s3URL {
		t.Errorf("expected S3URL %s, got %s", s3URL, video.S3URL)
	}
	if video.UpdatedAt.Before(video.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

func TestSetS3ManifestURL(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	s3ManifestURL := "https://s3.amazonaws.com/bucket/manifest.m3u8"
	video.SetS3ManifestURL(s3ManifestURL)

	if video.S3ManifestPath != s3ManifestURL {
		t.Errorf("expected S3ManifestPath %s, got %s", s3ManifestURL, video.S3ManifestPath)
	}
	if video.UpdatedAt.Before(video.CreatedAt) {
		t.Error("expected UpdatedAt to be after CreatedAt")
	}
}

func TestIsCompleted(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	video.MarkAsCompleted(video.HLSPath, video.ManifestPath)

	if !video.IsCompleted() {
		t.Error("expected video to be completed")
	}
}

func TestGenerateOutputPath(t *testing.T) {
	video := NewVideo("Test Video", "/path/to/video.mp4")
	baseDir := "/output"
	expectedOutputPath := filepath.Join(baseDir, "/converted/", video.ID)
	outputPath := video.GenerateOutputPath(baseDir)

	if outputPath != expectedOutputPath {
		t.Errorf("expected output path %s, got %s", expectedOutputPath, outputPath)
	}
}
