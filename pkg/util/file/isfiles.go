package file

import (
	"fmt"
	"os"
)

// is file there
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println("File does NOT exist.")
		return false
	} else if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
		return false
	} else {
		return true
	}
}
