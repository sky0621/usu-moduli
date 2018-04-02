package configs

import "github.com/spf13/viper"

// Config ...
type Config struct {
	TargetDir        string
	Outdir           string
	FilterOutProject string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		TargetDir:        viper.GetString("targetDir"),
		Outdir:           viper.GetString("outdir"),
		FilterOutProject: viper.GetString("filterOutProject"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
