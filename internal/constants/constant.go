package constants

import (
	"os"
	"path/filepath"
)

var (
	BaseURL    = "https://github.com"
	SyncURl    = "https://api.github.com/repos"
	LocalFile  = filepath.Join(os.Getenv("HOME"), ".local", "share", "lazydot", "local.yml")
	SyncFile   = filepath.Join(os.Getenv("HOME"), ".local", "share", "lazydot", "sync.yml")
	ConfigFile = filepath.Join(os.Getenv("HOME"), ".config", "lazydot", "config.yml")

	// ConfigFile       = "/mnt/e/projects/golang/lazydot/config/config.yml"
	DownloadLocation = filepath.Join(os.Getenv("HOME"), ".cache", "lazydot", "pkg")
)
