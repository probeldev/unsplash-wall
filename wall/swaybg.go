package wall

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type swaybg struct{}

func GetSwaybg() *swaybg {
	s := swaybg{}

	return &s
}

func (s *swaybg) getSwaybgPID() ([]int, error) {
	pids := []int{}

	cmd := exec.Command("pidof", "swaybg")
	output, err := cmd.Output()
	if err != nil {
		return pids, fmt.Errorf("swaybg не запущен или ошибка pgrep: %v", err)
	}

	// Преобразуем вывод в число (PID)
	pidStr := strings.TrimSpace(string(output))
	pidArr := strings.Split(pidStr, "\n")
	for _, p := range pidArr {
		pid, err := strconv.Atoi(p)
		if err != nil {
			return pids, fmt.Errorf("неверный PID: %v", err)
		}

		pids = append(pids, pid)
	}

	return pids, nil
}

func (s *swaybg) SetWallpaper(imagePath string) error {
	// Preload image
	// Проверяем, существует ли файл
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("Файл %s не найден!\n", imagePath)
		os.Exit(1)
	}

	// Получаем PID текущего swaybg
	oldPids, err := s.getSwaybgPID()
	if err != nil {
		fmt.Println("Предыдущий swaybg не найден (но запустим новый)", err)
	}

	// Запускаем новый swaybg
	cmd := exec.Command("swaybg", "-i", imagePath, "-m", "fill")
	err = cmd.Start()
	if err != nil {
		log.Println(err)
	}

	go func() {
		time.Sleep(10 * time.Second)
		for _, oldPID := range oldPids {
			// Если старый swaybg был найден, убиваем его
			if oldPID != 0 {
				fmt.Printf("Убиваем старый swaybg (PID: %d)\n", oldPID)
				err := syscall.Kill(oldPID, syscall.SIGTERM)
				if err != nil {
					fmt.Printf("Ошибка при убийстве процесса %d: %v\n", oldPID, err)
				}
			}
		}
	}()

	fmt.Printf("Обои установлены: %s\n", imagePath)
	return nil
}
