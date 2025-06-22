package model

// MetaData struct
type MetaData struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	URL     string `yaml:"url"`
	ConfDir string `yaml:"confdir"`
}
type Config struct {
	Username string `yaml:"username"`
	Project  string `yaml:"project"`
}

// BaseURL variable
// var BaseURL = "https://github.com/alicemarple/dotfiles/releases/download"

// https://github.com/alicemarple/dotfiles/releases/download/shell-v0.2.0/shell-v0.2.0.tar.gz
