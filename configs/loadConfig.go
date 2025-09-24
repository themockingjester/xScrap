package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Below functions reads the configuration from the config file
func LoadConfig(config *AppConfig) {

	// Read the YAML file
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Panicf("Error reading YAML file: %v\n", err)
		return
	}

	// Parse the YAML into the Config struct

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Panicf("Error unmarshaling YAML: %v\n", err)
		return
	}

}
