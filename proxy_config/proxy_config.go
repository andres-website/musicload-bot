package proxy_config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	UseProxy        bool   `yaml:"use_proxy"`
	Proxy           string `yaml:"proxy"`
	Use_youtube_api bool   `yaml:"use_youtube_api"`
	Youtube_api_key string `yaml:"youtube_api_key"`
}

var AppConfig *Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	AppConfig = &Config{}
	if err := yaml.Unmarshal(data, AppConfig); err != nil {
		log.Fatalf("error: %v", err)
	}
	// log.Println("AppConfig:")
	// log.Println(AppConfig)
}
