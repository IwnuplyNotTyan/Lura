package data

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Language     string   `toml:"language"`
	Score        int      `toml:"score"`
	Achievements []string `toml:"achievements"`
}

func GetConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "lura", "config.toml")
}

func TouchConfig(player *Player) Config {
	configPath := GetConfigPath()
	if configPath == "" {
		return Config{}
	}

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Printf("Error creating config directory: %v", err)
		return Config{}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		initialConfig := Config{
			Language:     "", 
			Score:        0, 
			Achievements: []string{},
		}

		if err := SaveConfig(configPath, initialConfig); err != nil {
			log.Printf("Error creating initial config: %v", err)
			return Config{}
		}
		return initialConfig
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return Config{}
	}
	return cfg
}

func LoadConfig(filename string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(filename, &cfg)
	return cfg, err
}

func SaveConfig(filename string, cfg Config) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return toml.NewEncoder(file).Encode(cfg)
}
