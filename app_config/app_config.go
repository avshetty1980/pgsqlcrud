package app_config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "config.yml"

type AppConfig struct {
	DB string `yaml:"db"`
	// Redis   string `yaml:"redis"`
	// FbOauth struct {
	// 	ClientID     string `yaml:"client_id"`
	// 	ClientSecret string `yaml:"client_secret"`
	// } `yaml:"fb_oauth"`
}

var Cfg AppConfig

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("could not open config file: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatalf("could not decode yaml file: %v", err)
	}

}
