package web

import (
	"fmt"
	"net/http"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"gopkg.in/yaml.v2"
)

func GetSyncData() {
	url := constants.SyncURl

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var meta []model.SyncData
	if err := yaml.NewDecoder(resp.Body).Decode(&meta); err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}

	var data []model.SyncData
	for _, rm := range meta {
		data = append(data, model.SyncData{
			Name:    rm.Name,
			Version: rm.Version,
			ConfDir: rm.ConfDir,
		})
	}

	err = file.WriteYml(constants.SyncFile, data)
	if err != nil {
		fmt.Println("Error writing sync.yml:", err)
		return
	}
}
