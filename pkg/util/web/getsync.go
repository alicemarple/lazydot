package web

import (
	"fmt"
	"net/http"

	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"gopkg.in/yaml.v2"
)

func GetSyncData() {
	url := "https://raw.githubusercontent.com/alicemarple/dotfiles/refs/heads/master/metdata.yml"

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

	// for _, m := range meta {
	// 	fmt.Printf("Name: %s, Version: %s, ConfDir: %s\n", m.Name, m.Version, m.ConfDir)
	// }
	var data []model.MetaData
	for _, rm := range meta {
		data = append(data, model.MetaData{
			Name:    rm.Name,
			Version: rm.Version,
			ConfDir: rm.ConfDir,
			URL:     "", // Or set this if you have a value
		})
	}

	// ✅ Call your WriteYml
	err = file.WriteYml("/mnt/e/projects/golang/lazydot/sync/sync.yml", data) // Adjust path
	if err != nil {
		fmt.Println("Error writing sync.yml:", err)
		return
	}

	fmt.Println("✅ sync.yml updated successfully.")
}
