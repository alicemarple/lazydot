package util

import (
	"fmt"
	"os"

	"github.com/alicemarple/lazydot/pkg/model"
	"gopkg.in/yaml.v2"
)

// ReadYml function
func ReadYml(filename string) []model.MetaData {
	isthere := FileExists(filename)
	if isthere {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		var md []model.MetaData
		if err := yaml.Unmarshal(data, &md); err != nil {
			fmt.Println(err)
			return nil
		}
		return md

	}
	return nil
}
