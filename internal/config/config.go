package config

import (
	"os"
	"encoding/json"
	"fmt"
	"errors"
)
type Config struct {
	Url     string `json:"db_url"`
}

func Read() (*Config, error) {
	file, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}
	file += "/.gatorconfig.json"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// if file does not exit, create a default Config
		defaultConfig := Config{Url: "http://localhost:8080"}
		data, err := json.Marshal(defaultConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal default config: %w", err)
		}
		if err := os.WriteFile(file, data, 0644); err != nil {
			return nil, fmt.Errorf("failed to write default config file: %w", err)
		}
		return &defaultConfig, nil
	}
	

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	if len(data) == 0 {
		return nil, errors.New("config file is empty")
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if config.Url == "" {
		return nil, errors.New("db_url is not set in the config file")
	}
	return &config, nil
	
}

func (c *Config) SetUser(user string) error {
	if user == "" {
		return errors.New("user cannot be empty")
	}
	c.Url = user
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	file, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	file += "/.gatorconfig.json"
	if err := os.WriteFile(file, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}
