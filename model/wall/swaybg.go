package wall

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type swaybg struct{}

func GetSwaybg() *swaybg {
	s := swaybg{}

	return &s
}

func (s *swaybg) getSwaybgPID() (int, error) {
	// Получаем PID swaybg через pgrep
	cmd := exec.Command("pgrep", "swaybg")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("swaybg не запущен или ошибка pgrep: %v", err)
	}

	// Преобразуем вывод в число (PID)
	pidStr := strings.TrimSpace(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, fmt.Errorf("неверный PID: %v", err)
	}

	return pid, nil
}

func (s *swaybg) SetWallpaper(imagePath string) error {
	// Preload image
	// Проверяем, существует ли файл
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("Файл %s не найден!\n", imagePath)
		os.Exit(1)
	}

	// Получаем PID текущего swaybg
	oldPID, err := s.getSwaybgPID()
	if err != nil {
		fmt.Println("Предыдущий swaybg не найден (но запустим новый)")
	}

	// Запускаем новый swaybg
	cmd := exec.Command("swaybg", "-i", imagePath, "-m", "fill")
	err = cmd.Start()
	if err != nil {
		log.Println(err)
	}

	// Если старый swaybg был найден, убиваем его
	if oldPID != 0 {
		fmt.Printf("Убиваем старый swaybg (PID: %d)\n", oldPID)
		err := syscall.Kill(oldPID, syscall.SIGTERM)
		if err != nil {
			fmt.Printf("Ошибка при убийстве процесса %d: %v\n", oldPID, err)
		}
	}

	fmt.Printf("Обои установлены: %s\n", imagePath)
	return nil
}
