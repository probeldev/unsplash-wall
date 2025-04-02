package wall

import (
	"fmt"
	"os/exec"
)

type hyprpaper struct{}

func GetHyprpaper() *hyprpaper {
	h := hyprpaper{}

	return &h
}

func (h *hyprpaper) SetWallpaper(imagePath string) error {
	// Preload image
	cmd := exec.Command("hyprctl", "hyprpaper", "preload", imagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to preload wallpaper: %v", err)
	}

	// Set wallpaper
	cmd = exec.Command("hyprctl", "hyprpaper", "wallpaper", ",", imagePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set wallpaper: %v", err)
	}

	return nil
}
