package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getPathSizeBytes(path, recursive, all)

	if err != nil {
		return "", err
	}
	return formatSize(size, human), nil
}

func getPathSizeBytes(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		return 0, err
	}

	var totalSize int64

	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		info, err := entry.Info()

		if err != nil {
			return 0, err
		}

		if !entry.IsDir() {
			totalSize += info.Size()
			continue
		}

		if !recursive {
			continue
		}

		nestedPath := filepath.Join(path, entry.Name())

		nestedSize, err := getPathSizeBytes(nestedPath, recursive, all)

		if err != nil {
			return 0, err
		}

		totalSize += nestedSize
	}
	return totalSize, nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	unit := 0

	for value >= 1024 && unit < len(units)-1 {
		value /= 1024
		unit++
	}

	if unit == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unit])
}
