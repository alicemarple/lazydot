package file

import (
	"fmt"
	"os"
)

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	fmt.Println("File created successfully:", filename)
	return nil
}
