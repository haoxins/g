package g

import (
	"os"
)

// Exists checks whether a file or directory exists.
func Exists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CatFile reads a file and returns its content as string.
func CatFile(filepath string) string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return ""
	}
	return string(bytes)
}
