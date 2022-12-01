package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) (string, error) {
	inputFile, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("os.ReadFile: %w", err)
	}
	return strings.Trim(string(inputFile), "\r\n"), nil
}
