package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// GetConfig reads the standard config file and returns value of key
func GetConfig(key string) (string, error) {
	homeDir := os.Getenv("HOME")
	kciConfigPath := path.Join(homeDir, ".kci", "config.yaml")

	contents, err := ioutil.ReadFile(kciConfigPath)
	if err != nil {
		return "", err
	}

	config := make(map[string]string)
	yaml.Unmarshal(contents, &config)

	return config[key], nil
}
