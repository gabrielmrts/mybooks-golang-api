package utils

import (
	"fmt"
	"os"
)

func GetTemplatesFolder() string {
	path, _ := os.Getwd()
	path = fmt.Sprintf("%s/%s", path, "internal/templates")

	return path
}
