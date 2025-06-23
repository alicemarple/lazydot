package file

import (
	"fmt"
	"os"
)

func IsEmpty(filename string) bool {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return false
	}
	defer f.Close()

	buf := make([]byte, 1)
	n, _ := f.Read(buf)

	if n == 0 {
		return true
	} else {
		fmt.Println("file is not empty")
		return false
	}
}
