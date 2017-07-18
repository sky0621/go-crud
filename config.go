package go_crud

import "github.com/spf13/viper"

// Config ...
type Config struct {
	Target string
	Tables string
	Filter *FilterConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Target: viper.GetString("target"),
		Tables: viper.GetString("tables"),
		Filter: NewFilterConfig(),
	}
}

// FilterConfig ...
type FilterConfig struct {
	OutDir  []string
	OutFile []string
	InDir   []string
	InFile  []string
}

// NewFilterConfig ...
func NewFilterConfig() *FilterConfig {
	return &FilterConfig{
		OutDir:  viper.GetStringSlice("filter.outdir"),
		OutFile: viper.GetStringSlice("filter.outfile"),
		InDir:   viper.GetStringSlice("filter.indir"),
		InFile:  viper.GetStringSlice("filter.infile"),
	}
}

// ReadConfig ...
func ReadConfig(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	return viper.ReadInConfig()
}
