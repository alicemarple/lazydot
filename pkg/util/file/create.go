package file

import (
	"fmt"
	"os"

	"github.com/alicemarple/lazydot/internal/constants"
)

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	fmt.Println(constants.SuccessStyle.Render(fmt.Sprint("File created successfully:")), filename)
	return nil
}
