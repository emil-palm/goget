package config;

import (
	"testing"
	"github.com/spf13/viper"
)

func init() {
	// Add the resources folder to the load path of viper so we can parse the example config
	viper.AddConfigPath("resources/")
}		

func Test_LoadConfig(t *testing.T) {
	cfg,err := GetConfig()
	if err != nil {
		t.Error(err)
	}

	if len(cfg.RSSParser) == 0 {
		t.Error("Config seams to have failed somewhere")
	}
}