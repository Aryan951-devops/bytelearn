package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

func IsImageAllowed(fileHeader *multipart.FileHeader) error {

	// Validate file size
	if fileHeader.Size > 10<<20 {
		return errors.New("image size should be less than 10MB")
	}

	// Validate extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExtensions[ext] {
		return errors.New("invalid image extension")
	}

	// Open file stream
	file, err := fileHeader.Open()

	if err != nil {
		return errors.New("failed to open uploaded file")
	}

	defer file.Close()

	// Read first 512 bytes
	buffer := make([]byte, 512)

	_, err = file.Read(buffer)

	if err != nil {
		return errors.New("failed to read uploaded file")
	}

	// Reset cursor
	_, err = file.Seek(0, io.SeekStart)

	if err != nil {
		return errors.New("failed to reset file cursor")
	}

	// Detect MIME type
	contentType := http.DetectContentType(buffer)

	allowedMimeTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}

	if !allowedMimeTypes[contentType] {
		return errors.New("invalid image mime type")
	}

	return nil
}
