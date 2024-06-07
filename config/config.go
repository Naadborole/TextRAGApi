package config

import (
	"encoding/json"
	"os"
)

func getConfigValue(key string) (string, error) {
	configFile, err := os.ReadFile("../config.json")
	if err != nil {
		return "", err
	}
	var config map[string]string
	json.Unmarshal(configFile, &config)
	return config[key], nil
}
