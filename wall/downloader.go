package wall

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"math/rand"
)

type wallDownloader struct{}

func GetWallDownloader() *wallDownloader {
	w := wallDownloader{}

	return &w
}

func (w *wallDownloader) downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed download image")
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (w *wallDownloader) DownloadRandomImage(
	downloadURLs []string,
	tmpDir string,
	timestamp int64,
) (
	string,
	error,
) {
	selectedURL := downloadURLs[rand.Intn(len(downloadURLs))]

	imagePath := filepath.Join(tmpDir, fmt.Sprintf("wall-%d.jpg", timestamp))
	err := w.downloadFile(selectedURL, imagePath)
	if err != nil {
		return "", err
	}
	return imagePath, nil
}
