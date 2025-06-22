package util

import (
	"fmt"
	"os"
)

// is file there
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Printf("Error checking file %s: %v\n", filename, err)
		return false
	}
	return !info.IsDir()
}
