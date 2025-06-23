package commmand

import (
	"fmt"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/web"
)

// TODO: Add functionality to extract the .tar.gz and put it config dir
// sync function
func Sync(filename string, packageName string) {
	md := file.ReadYml[model.MetaData](filename)

	found := false

	for _, v := range md {
		if packageName == v.Name {
			found = true
			fmt.Println(v.Name, v.ConfDir, v.Version)

			// check package info from sync.yml
			newURL := fmt.Sprintf("%s/%s-%s/%s-%s.tar.gz", constants.BaseURL, v.Name, v.Version, v.Name, v.Version)
			fileLoaction := fmt.Sprintf("%s%s-%s.tar.gz", constants.DownloadLocation, v.Name, v.Version)
			newFileName := fmt.Sprintf("%s-%s.tar.gz", v.Name, v.Version)
			newConfigDir := fmt.Sprintf("~/.config/%s", v.Name)

			// download file
			err := web.DownloadFromURL(newURL, fileLoaction)
			fmt.Println(newURL)
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
			file.WriteYml("/mnt/e/projects/golang/lazydot/local/local.yml", lm)

			// get sync database
			// util.GetSyncData()

			// extract file

			// put to its config dir
		}
	}
	if !found {
		fmt.Println("first update the database by using -y")
	}
}
