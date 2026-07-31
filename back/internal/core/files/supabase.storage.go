package files

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type supabaseImageStorage struct {
	url    string
	key    string
	bucket string
}

func NewSupabaseImageStorage(
	bucket string,
	url string,
	key string,
) ImageStorage {
	return &supabaseImageStorage{
		url:    url,
		key:    key,
		bucket: bucket,
	}
}

func (s *supabaseImageStorage) SaveImage(
	file multipart.File,
	header *multipart.FileHeader,
	folder string,
	prefix string,
	ownerID int,
) (string, error) {

	ext := strings.ToLower(filepath.Ext(header.Filename))

	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	const MaxImageSize = 5 << 20 // 5MB

	if header.Size > MaxImageSize {
		return "", fmt.Errorf("imagen demasiado grande")
	}

	if !allowed[ext] {
		return "", fmt.Errorf("formato de imagen invalido")
	}

	filename := fmt.Sprintf(
		"%s_%d_%d%s",
		prefix,
		ownerID,
		time.Now().UnixNano(),
		ext,
	)

	path := strings.Trim(folder, "/") + "/" + filename

	body, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf(
		"%s/storage/v1/object/%s/%s",
		s.url,
		s.bucket,
		path,
	)

	req, err := http.NewRequest(
		http.MethodPost,
		endpoint,
		bytes.NewReader(body),
	)

	if err != nil {
		return "", err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+s.key,
	)

	req.Header.Set(
		"apikey",
		s.key,
	)

	req.Header.Set(
		"Content-Type",
		header.Header.Get("Content-Type"),
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		responseBody, _ := io.ReadAll(resp.Body)

		return "",
			fmt.Errorf(
				"supabase upload error: %s",
				string(responseBody),
			)
	}

	publicURL := fmt.Sprintf(
		"%s/storage/v1/object/public/%s/%s",
		s.url,
		s.bucket,
		path,
	)

	return publicURL, nil
}
