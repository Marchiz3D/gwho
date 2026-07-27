package cmd

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Profiles   map[string]Profile `yaml:"profiles"`
	Workspaces map[string]string  `yaml:"workspaces"`
}

type Profile struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func LoadConfig() (Config, error) {
	configPath, err := getConfigFile()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{
				Profiles:   make(map[string]Profile),
				Workspaces: make(map[string]string),
			}, nil
		}
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	if config.Profiles == nil {
		config.Profiles = make(map[string]Profile)
	}
	if config.Workspaces == nil {
		config.Workspaces = make(map[string]string)
	}

	return config, nil
}

func SaveConfig(config Config) error {
	configPath, err := getConfigFile()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

func getConfigFile() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(baseDir, "gwho", "config.yaml")

	return configPath, nil
}
