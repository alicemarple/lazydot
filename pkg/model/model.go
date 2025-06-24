package model

// interface
// Data
type Data interface {
	GetName() string
}

func (m MetaData) GetName() string { return m.Name }
func (c Config) GetName() string   { return c.Name }
func (s SyncData) GetName() string { return s.Name }

type PrintData interface {
	GetPName() string
	GetPVersion() string
	GetPConfDir() string
}

// methods for MetaData
func (m MetaData) GetPName() string    { return m.Name }
func (m MetaData) GetPVersion() string { return m.Version }
func (m MetaData) GetPConfDir() string { return m.ConfDir }

// methods for SyncData
func (s SyncData) GetPName() string    { return s.Name }
func (s SyncData) GetPVersion() string { return s.Version }
func (s SyncData) GetPConfDir() string { return s.ConfDir }

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
