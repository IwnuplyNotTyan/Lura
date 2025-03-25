package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	lua "github.com/yuin/gopher-lua"
)

type Config struct {
	Language     string   `toml:"language"`
	Score        int      `toml:"score"`
	Achievements []string `toml:"achievements"`
}

var (
	debugMode       = flag.Bool("debug", false, "Enable debug shell")
	specificMonster *Monster
)

func main() {
	flag.Parse()
	L := lua.NewState()
	defer L.Close()
	clearScreen()
	dialWelcome()
	fmt.Printf("\n")

	// Initialize config first to get language
	player := Player{}     // Temporary empty player for config initialization
	cfg := config(&player) // This will load or create config with language

	// Set language from config
	lang = cfg.Language
	if lang == "" {
		lang = selectLanguage()
		cfg.Language = lang
		saveConfig(getConfigPath(), cfg) // Save the selected language
	}

	seedData()

	registerTypes(L)
	L.SetGlobal("lang", lua.LString(lang))

	if err := AutoLoadMods(L); err != nil {
		log.Fatalf("Failed to auto-load mods: %v", err)
	}

	weaponType, weaponDamage := getRandomWeapon()
	player = Player{
		WeaponType: weaponType,
		Damage:     weaponDamage * rng(),
		HP:         100,
		maxHP:      100,
		Coins:      0,
		Stamina:    100,
		maxStamina: 100,
		heart:      1,
		loc:        1,
	}

	if *debugMode {
		DebugShell(L, &player)
	}
	cfg = config(&player)
	fight(&player, specificMonster, &cfg)
}

func selectLanguage() string {
	var selectedLang string
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(" Select Language").
				Options(
					huh.NewOption("English", "en"),
					huh.NewOption("Українська", "ua"),
					huh.NewOption("Беларускiй", "be"),
				).
				Value(&selectedLang),
		),
	)

	if err := f.Run(); err != nil {
		log.Printf("Error selecting language: %v", err)
		return "en"
	}

	switch selectedLang {
	case "en", "ua", "be":
		return selectedLang
	default:
		return "en"
	}
}

func getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "lura", "config.toml")
}

func config(player *Player) Config {
	configPath := getConfigPath()
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
			Language:     lang,
			Score:        player.score,
			Achievements: []string{},
		}

		if err := saveConfig(configPath, initialConfig); err != nil {
			log.Printf("Error creating initial config: %v", err)
			return Config{}
		}
		return initialConfig
	}

	cfg, err := loadConfig(configPath)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return Config{}
	}
	return cfg
}

func loadConfig(filename string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(filename, &cfg)
	return cfg, err
}

func saveConfig(filename string, cfg Config) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return toml.NewEncoder(file).Encode(cfg)
}
