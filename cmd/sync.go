package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: Make this code modular
type metaData struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	ConfDir string `yaml:"confdir"`
}
type localMetaData struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Url       string `yaml:"url"`
	ConfigDir string `yaml:"confdir"`
}

var baseURL string = "https://github.com/alicemarple/dotfiles/releases/download"

// https://github.com/alicemarple/dotfiles/releases/download/shell-v0.2.0/shell-v0.2.0.tar.gz

// is file there
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		fmt.Printf("Error checking file %s: %v\n", filename, err)
		return false
	}
	return !info.IsDir()
}

// read from the yml (sync.yml)
func readYml(filename string, packageName string) {
	isthere := fileExists(filename)
	if isthere {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var md []metaData
		if err := yaml.Unmarshal(data, &md); err != nil {
			panic(err)
		}

		for _, v := range md {
			if packageName == v.Name {
				fmt.Println(v.Name, v.ConfDir, v.Version)

				// check package info from sync.yml
				newURL := fmt.Sprintf("%s/%s-v%s/%s-v%s.tar.gz", baseURL, v.Name, v.Version, v.Name, v.Version)
				newFileName := fmt.Sprintf("%s-v%s.tar.gz", v.Name, v.Version)

				// download file
				err := downloadFromURL(newURL, newFileName)
				if err != nil {
					fmt.Println("error:%v", err)
				}
				// update the local.yml according to downloaded file

				lm := []localMetaData{
					{
						Name:      newFileName,
						Version:   v.Version,
						Url:       newURL,
						ConfigDir: "~/.config/xyz",
					},
				}
				writeYml("/mnt/e/projects/golang/lazydot/local/local.yml", lm)
				// fmt.Print(lm)
			}
		}

	} else {
		fmt.Println("fileNotexists")
	}
}

func writeYml(filename string, data interface{}) error {
	// FIX: It can have multiple entry for same package due to appending into file
	// TODO: Do find and replace of the text
	out, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// 2️⃣ Open the file in Append Mode
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 3️⃣ Append the data
	_, err = file.Write(out)
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}
	return nil
}

// download using the newURL

func downloadFromURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error during GET request: %w", err)
	}
	defer resp.Body.Close()

	// 2️⃣ Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 3️⃣ Create the destination file
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// 4️⃣ Copy response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error while downloading: %w", err)
	}

	fmt.Printf("File saved to %s\n", filename)
	return nil
}

// update the local.yml file
// func (string name, string version, string url, string configDir)  {
//
//
// }
