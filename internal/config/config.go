package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func Load(p string) error {
	configName := strings.TrimRight(filepath.Base(p), filepath.Ext(p))
	configPath := strings.TrimRight(p, configName)

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	return nil
}
