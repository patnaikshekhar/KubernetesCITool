package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Config creates a .kci directory in home and stores configuration
func AddConfig(key string, value string) {
	// Check if file for kci exists
	homeDir := os.Getenv("HOME")
	kciConfigPath := path.Join(homeDir, ".kci", "config.yaml")
	if _, err := os.Stat(kciConfigPath); err != nil {
		// Create Directory
		if err := os.MkdirAll(path.Join(homeDir, ".kci"), 0700); err != nil {
			log.Fatalf("Could not create config directory: %s", err.Error())
		}

		// Create File
		ioutil.WriteFile(kciConfigPath, []byte{}, 0644)
	}

	contents, err := ioutil.ReadFile(kciConfigPath)
	if err != nil {
		log.Fatalf("Error when reading config: %s", err.Error())
	}

	var config map[string]string
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		log.Fatalf("Error when reading config: %s", err.Error())
	}

	if config == nil {
		config = make(map[string]string)
	}

	// Add key value pair
	config[key] = value

	newContents, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Error when writing config: %s", err.Error())
	}

	err = ioutil.WriteFile(kciConfigPath, newContents, 0644)
	if err != nil {
		log.Fatalf("Error when writing config: %s", err.Error())
	}
}
