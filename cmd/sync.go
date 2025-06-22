package cmd

import (
	"fmt"

	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util"
)

// TODO: Add functionality to extract the .tar.gz and put it config dir
// sync function
func sync(filename string, packageName string) {
	var md []model.MetaData
	md = util.ReadYml(filename)
	found := false

	for _, v := range md {
		if packageName == v.Name {
			found = true
			fmt.Println(v.Name, v.ConfDir, v.Version)

			// check package info from sync.yml
			newURL := fmt.Sprintf("%s/%s-v%s/%s-v%s.tar.gz", model.BaseURL, v.Name, v.Version, v.Name, v.Version)
			newFileName := fmt.Sprintf("%s-v%s.tar.gz", v.Name, v.Version)
			newConfigDir := fmt.Sprintf("~/.config/%s", v.Name)

			// download file
			err := util.DownloadFromURL(newURL, newFileName)
			if err != nil {
				fmt.Printf("error:%v", err)
			}

			// update the local.yml according to downloaded file
			lm := []model.MetaData{
				{
					Name:    newFileName,
					Version: v.Version,
					URL:     newURL,
					ConfDir: newConfigDir,
				},
			}
			util.WriteYml("/mnt/e/projects/golang/lazydot/local/local.yml", lm)

			// extract file

			// put to its config dir
		}
	}
	if !found {
		fmt.Println("first update the database by using -Y")
	}
}
