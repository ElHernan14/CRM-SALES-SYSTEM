package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageStorage interface {
	SaveImage(file multipart.File, header *multipart.FileHeader, folder string, prefix string, ownerID int) (string, error)
}

type localImageStorage struct {
	uploadDir string
}

func NewLocalImageStorage(uploadDir string) ImageStorage {
	return &localImageStorage{uploadDir: uploadDir}
}

func (s *localImageStorage) SaveImage(file multipart.File, header *multipart.FileHeader, folder string, prefix string, ownerID int) (string, error) {
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		return "", fmt.Errorf("formato de imagen invalido")
	}

	absoluteDir := filepath.Join(s.uploadDir, folder)
	if err := os.MkdirAll(absoluteDir, 0755); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s_%d_%d%s", prefix, ownerID, time.Now().UnixNano(), ext)
	absolutePath := filepath.Join(absoluteDir, filename)

	dst, err := os.Create(absolutePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	publicFolder := strings.ReplaceAll(folder, string(os.PathSeparator), "/")
	return "/uploads/" + strings.Trim(publicFolder, "/") + "/" + filename, nil
}
