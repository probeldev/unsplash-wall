package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/probeldev/fastlauncher/imageHelper"
	"github.com/probeldev/fastlauncher/model"
	"github.com/probeldev/fastlauncher/unsplash"
	"github.com/probeldev/fastlauncher/wall"
)

func run() {
	// Создаем временную директорию
	tmpDir := "/tmp/wall"
	err := os.MkdirAll(tmpDir, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Получаем текущую метку времени
	timestamp := time.Now().Unix()

	unsplashParser := unsplash.GetUnsplashParser()
	jsonData := unsplashParser.GetJsonData()

	// Парсим JSON
	var data model.Root
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Собираем все URL для скачивания
	var downloadURLs []string

	photos := data.ReduxInitialState.Entities.Photos
	for _, p := range photos {
		downloadURLs = append(downloadURLs, p.Links.Download)
	}

	if len(downloadURLs) == 0 {
		log.Fatal("No download URLs found")
	}

	// Удаляем старые изображения
	files, err := filepath.Glob(filepath.Join(tmpDir, "*.jpg"))
	if err != nil {
		log.Printf("Warning: failed to list old jpg files: %v", err)
	}
	for _, f := range files {
		os.Remove(f)
	}

	imagePath := ""
	countTry := 0
	for {
		imagePath, err = downloadRandomImage(downloadURLs, tmpDir, timestamp)
		if err != nil {
			log.Println(err)
		}

		ih := imageHelper.GetImageHelper()
		if ih.IsHorizontal(imagePath) {
			break
		}

		countTry++
		time.Sleep(5 * time.Second)
		log.Println("sleep")

		if countTry > 10 {
			log.Panic("failde download")
		}
	}

	log.Println(imagePath)

	// Устанавливаем обои
	err = wall.GetSwaybg().SetWallpaper(imagePath)
	if err != nil {
		log.Fatalf("Failed to set wallpaper: %v", err)
	}

	fmt.Println("Wallpaper set successfully!")
}

func downloadFile(url, filepath string) error {
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

// Вспомогательная функция для поиска подстроки
func stringIndex(s, substr string) int {
	idx := len(substr) + strings.Index(s, substr)
	if idx < len(substr) {
		return -1
	}
	return idx
}

func downloadRandomImage(
	downloadURLs []string,
	tmpDir string,
	timestamp int64,
) (
	string,
	error,
) {
	// Выбираем случайный URL
	selectedURL := downloadURLs[rand.Intn(len(downloadURLs))]

	// Скачиваем изображение
	imagePath := filepath.Join(tmpDir, fmt.Sprintf("wall-%d.jpg", timestamp))
	err := downloadFile(selectedURL, imagePath)
	if err != nil {
		return "", err
	}
	return imagePath, nil
}

func main() {

	for {
		run()
		time.Sleep(30 * time.Minute)
	}
}
