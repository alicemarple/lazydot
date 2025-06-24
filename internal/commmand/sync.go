package commmand

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alicemarple/lazydot/internal/constants"
	"github.com/alicemarple/lazydot/pkg/model"
	"github.com/alicemarple/lazydot/pkg/util/file"
	"github.com/alicemarple/lazydot/pkg/util/web"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// TODO: Add functionality to extract the .tar.gz and put it config dir
// sync function
func Sync(filename string, packageName string) {
	md := file.ReadYml[model.MetaData](filename)

	found := false

	for _, v := range md {
		if packageName == v.Name {
			found = true
			printSyncTable([]model.MetaData{v})
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
			Runscript(packageName)
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
	// cmd := exec.Command(scriptPath)
	cmd := exec.Command("/bin/sh", scriptPath, packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run script: %w", err)
	}
	return nil
}

func Runscript(packageName string) {
	script := "/mnt/e/projects/golang/lazydot/scripts/place.sh"

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

func printSyncTable(dt []model.MetaData) {
	// Styles
	var (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("241")

		headerStyle  = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(23)
		oddRowStyle  = cellStyle.Foreground(gray)
		evenRowStyle = cellStyle.Foreground(lightGray)
	)

	// Build table
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(purple)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("Name", "Version", "Destination")

	// Add rows
	for _, v := range dt {
		t.Row(v.Name, v.Version, v.ConfDir)
	}

	// Print the table
	fmt.Println(t)
}
