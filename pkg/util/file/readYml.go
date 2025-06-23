package file

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// ReadYml function
func ReadYml[T any](filename string) []T {
	isthere := FileExists(filename)
	if isthere {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		var md []T
		if err := yaml.Unmarshal(data, &md); err != nil {
			fmt.Println(err)
			return nil
		}
		return md

	}
	return nil
}
