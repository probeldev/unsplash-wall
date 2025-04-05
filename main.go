package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/probeldev/fastlauncher/imageHelper"
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
	downloadURLs, err := unsplashParser.GetImageUrls()
	if err != nil {
		log.Println(err)
		// TODO возможно стоит как то обрабатывтаь????
		return
	}

	// Удаляем старые изображения
	files, err := filepath.Glob(filepath.Join(tmpDir, "*.jpg"))
	if err != nil {
		log.Printf("Warning: failed to list old jpg files: %v", err)
	}
	for _, f := range files {
		os.Remove(f)
	}

	downloader := wall.GetWallDownloader()

	imagePath := ""
	countTry := 0
	for {
		imagePath, err = downloader.DownloadRandomImage(downloadURLs, tmpDir, timestamp)
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

func main() {

	for {
		run()
		time.Sleep(30 * time.Minute)
	}
}
