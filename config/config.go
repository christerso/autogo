package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

const dbFile = ".autojump.json"

type Config struct {
	Paths map[string]int `json:"paths"`
}

func LoadConfig() (*Config, error) {
	configPath := getConfigPath()
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Paths: make(map[string]int)}, nil
		}
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing file")
		}
	}(file)

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	configPath := getConfigPath()
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing file")
		}
	}(file)

	encoder := json.NewEncoder(file)
	return encoder.Encode(config)
}

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting home directory")
	}
	return filepath.Join(home, dbFile)
}
