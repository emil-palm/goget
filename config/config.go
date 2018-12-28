package config;

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
)

type Config struct {
	RSSParser map[string]Section
	TorrentClient TorrentClient
	Logging struct {
		Level log.Level
	}
}

func init() {
	viper.SetConfigName("goget")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("$HOME/.config/")
	viper.AddConfigPath(".")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", e.Name)
	})

	go viper.WatchConfig()
}

func GetConfig() (*Config, error) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		return nil, err
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

/*

release:
- release_name
- urls
- time
- expiration_time
- score
- state
- flags

*******************************
Configuration;
sections:
	tv:
		urls:
			- www.xxxx.se/rss
			- www.yyyy.se/rss
		scoring:
			hebsub: 
				matching: title
				all: reject
			internal:
				matching: flag
				first: 0.8
				other: 0.01
		path: /mnt/tv/{title}/{season}
		matching:
			list:
				- Law and Order
			externalrss: http://www.trakt.tv
	movies:





*/