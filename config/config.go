package config

import (
	"encoding/json"
	"os"
)

func GetConfigValue(key string) string {
	configFile, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	var config map[string]string
	json.Unmarshal(configFile, &config)
	return config[key]
}
