package commmand

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/tables"
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
			tables.PrintTable([]model.MetaData{v})
			fmt.Printf("\n \n")

			// check package info from sync.yml
			sourceURL := fmt.Sprintf("%s/%s-%s/%s-%s.tar.gz", constants.BaseURL, v.Name, v.Version, v.Name, v.Version)
			fileLocation := fmt.Sprintf("%s/%s-%s.tar.gz", constants.DownloadLocation, v.Name, v.Version)

			// download file
			err := web.DownloadFromURL(sourceURL, fileLocation)
			fmt.Println(constants.InfoStyle.Render(fmt.Sprint("sourceURL:")), sourceURL)

			if err != nil {
				fmt.Printf("error:%v", err)
			}

			// update the local.yml according to downloaded file
			lm := []model.MetaData{
				{
					Name:    v.Name,
					Version: v.Version,
					URL:     sourceURL,
					ConfDir: v.ConfDir,
				},
			}
			file.WriteYml(constants.LocalFile, lm)

			// extract file
			// put to its config dir
			fmt.Println(constants.SuccessStyle.Render(fmt.Sprint("Extracting and Placing the files : ")))
			script := "/mnt/e/projects/golang/lazydot/scripts/place.sh"
			Runscript(packageName, script)
		}
	}
	if !found {
		fmt.Println("first update the database by using -y")
	}
}

func makeExecutable(scriptPath string) error {
	err := os.Chmod(scriptPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to chmod script: %w", err)
	}
	return nil
}

func runScript(scriptPath string, packageName string) error {
	cmd := exec.Command("/bin/sh", scriptPath, packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run script: %w", err)
	}
	return nil
}

func Runscript(packageName string, script string) {
	// 1️⃣ Make script executable
	if err := makeExecutable(script); err != nil {
		fmt.Println("Error making script executable:", err)
		return
	}

	// 2️⃣ Run the script
	if err := runScript(script, packageName); err != nil {
		fmt.Println("Error running script:", err)
		return
	}
}
