package api

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	URL      string `yaml:"URL"`
}

var AppConfig Config

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "subtui", "config.yaml"), nil
}

func LoadConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return fmt.Errorf("could not decode config: %v", err)
	}

	return nil
}

func SaveConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	return encoder.Encode(&AppConfig)
}
