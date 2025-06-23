package model

type Data interface {
	GetName() string
}

// MetaData struct
type MetaData struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	URL     string `yaml:"url"`
	ConfDir string `yaml:"confdir"`
}

type Config struct {
	Name    string `yaml:"username"`
	Project string `yaml:"project"`
}

type SyncData struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	ConfDir string `yaml:"confdir"`
}

func (m MetaData) GetName() string { return m.Name }
func (c Config) GetName() string   { return c.Name }
func (s SyncData) GetName() string { return s.Name }
