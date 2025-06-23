package setup

import (
	"fmt"
	"os"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/web"
	"gopkg.in/yaml.v2"
)

// read config.yml file
func ReadYmlSingle[T any](filename string) (T, error) {
	var item T
	isthere := file.FileExists(filename)
	if !isthere {
		return item, fmt.Errorf("file does not exist: %s", filename)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return item, fmt.Errorf("failed to read file: %w", err)
	}
	if err := yaml.Unmarshal(data, &item); err != nil {
		return item, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}
	return item, nil
}

// setup function
func Setup(filename string) {
	cfg, err := ReadYmlSingle[model.Config](filename)
	if err != nil {
		fmt.Println("config files read failed", err)
	}

	// setup urls
	constants.BaseURL = fmt.Sprintf("%s/%s/%s/releases/download", constants.BaseURL, cfg.Name, cfg.Project)
	constants.SyncURl = fmt.Sprintf("%s/%s/%s/releases", constants.SyncURl, cfg.Name, cfg.Project)

	// setup sync.yml
	isthere := file.FileExists(constants.SyncFile)
	istherelocal := file.FileExists(constants.LoaclFile)

	if !istherelocal {
		fmt.Println("Local file doesn't exist, creating it...")
		file.CreateFile(constants.LoaclFile)
	}
	if !isthere {
		fmt.Println("Sync file doesn't exist, creating it...")
		file.CreateFile(constants.SyncFile)
		web.GetSyncData()
	} else {
		isempty := file.IsEmpty(constants.SyncFile)
		if isempty {
			fmt.Println("File is empty, doing the needed work...")
			web.GetSyncData()
		} else {
			fmt.Println("File exists and is NOT empty, skipping...")
		}
	}
}
