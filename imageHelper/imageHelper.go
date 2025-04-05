package imageHelper

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type imageHelper struct{}

func GetImageHelper() *imageHelper {
	i := imageHelper{}

	return &i
}

func (i *imageHelper) IsHorizontal(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false // или обработать ошибку по-другому
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false // или обработать ошибку по-другому
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	return width > height
}
