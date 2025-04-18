package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/charmbracelet/log"
	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	"github.com/muesli/termenv"
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

//go:embed assets/*
var assetsFS embed.FS

func main() {
	flag.Parse()
	L := lua.NewState()
	defer L.Close()
	clearScreen()

	dialWelcome()

	player := Player{}
	cfg := config(&player)
	lang = cfg.Language
	if lang == "" {
		lang = selectLanguage()
		cfg.Language = lang
		saveConfig(getConfigPath(), cfg)
	}

	seedData()
	registerTypes(L)
	L.SetGlobal("lang", lua.LString(lang))

	if err := AutoLoadMods(L); err != nil {
		log.Fatalf("Failed to auto-load mods: %v", err)
	}

	//if loadedMods == nil {
	//	fmt.Print(termenv.String("  ").Foreground(termenv.ANSIWhite).Bold())
	//}

	//fmt.Print(termenv.String("  ").Foreground(termenv.ANSIMagenta).Bold())
	fmt.Printf("\n")

	weaponType, weaponDamage := getRandomWeapon()
	player = Player{
		WeaponType: weaponType,
		Damage:     weaponDamage * rng(),
		HP:         rngHp(),
		maxHP:      player.HP,
		Coins:      0,
		Stamina:    rngHp(),
		maxStamina: player.Stamina,
		heart:      1,
		loc:        1,
		monster:    false,
		Position:   0,
		Inventory: Inventory{
			Items:  make([]Item, 0),
			NextID: 1,
		},
	}

	if *debugMode {
		DebugShell(L, &player)
	}
	cfg = config(&player)
	fight(&player, specificMonster, &cfg, &Weapon{})
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

func (inv *Inventory) AddItem(name string, effect string, value int, price int) {
    // Check if item already exists
    for i, item := range inv.Items {
        if item.Name == name {
            inv.Items[i].Quantity++
            return
        }
    }
    
    // Add new item
    newItem := Item{
        ID:       inv.NextID,
        Name:     name,
        Quantity: 1,
        Effect:   effect,
        Value:    value,
        Price:    price,
    }
    inv.Items = append(inv.Items, newItem)
    inv.NextID++
}

func (inv *Inventory) RemoveItem(id int, quantity int) bool {
    for i, item := range inv.Items {
        if item.ID == id {
            if item.Quantity > quantity {
                inv.Items[i].Quantity -= quantity
            } else {
                // Remove item completely if quantity is 0 or negative
                inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
            }
            return true
        }
    }
    return false
}

func (player *Player) UseItem(id int) {
    for _, item := range player.Inventory.Items {
        if item.ID == id {
            switch item.Effect {
	    case "Material":
		    if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  You get %s, +%d material!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  Ви отримали %s, +%d матеріалів!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Вы атрымалі %s, +%d матэрыялаў!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    }
            case "heal":
                player.HP = min(player.HP+item.Value, player.maxHP)
                if lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, healed %d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "damage_boost":
                player.Damage += item.Value
                if lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, +%d damage for this fight!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d пошкоджень у цьому бою!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d пашкоджанняў у гэтым баі!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "stamina_restore":
                player.Stamina = min(player.Stamina+item.Value, player.maxStamina)
                if lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, restored %d stamina!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d витривалостi!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d вынослівасці!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
            }
            
            player.Inventory.RemoveItem(id, 1)
            return
        }
    }
    
    if lang == "en" {
        fmt.Println(termenv.String("  Item not found!").Foreground(termenv.ANSIRed))
    } else if lang == "ua" {
        fmt.Println(termenv.String("  Предмет не знайдено!").Foreground(termenv.ANSIRed))
    } else if lang == "be" {
        fmt.Println(termenv.String("  Прадмет не знойдзены!").Foreground(termenv.ANSIRed))
    }
}

func ShowInventory(player *Player) {
    if len(player.Inventory.Items) == 0 {
        if lang == "en" {
            fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
        } else if lang == "ua" {
            fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
        } else if lang == "be" {
            fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
        }
        return
    }
    
    if lang == "en" {
        fmt.Println(termenv.String("  Your inventory:").Foreground(termenv.ANSIMagenta).Bold())
    } else if lang == "ua" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    } else if lang == "be" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    }
    
    for _, item := range player.Inventory.Items {
        fmt.Printf(" [%d] %s x%d\n", item.ID, item.Name, item.Quantity)
    }
}

func UseItemMenu(player *Player) {
    if len(player.Inventory.Items) == 0 {
        if lang == "en" {
            fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
        } else if lang == "ua" {
            fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
        } else if lang == "be" {
            fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
        }
        return
    }
    
    var selectedID int
    options := make([]huh.Option[int], 0, len(player.Inventory.Items)+1)
    
    // Add "Cancel" option
    if lang == "en" {
        options = append(options, huh.NewOption("Cancel", -1))
    } else if lang == "ua" {
        options = append(options, huh.NewOption("Скасувати", -1))
    } else if lang == "be" {
        options = append(options, huh.NewOption("Адмяніць", -1))
    }
    
    // Add item options
    for _, item := range player.Inventory.Items {
        var desc string
        if lang == "en" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescription(item.Effect), item.Value)
        } else if lang == "ua" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionUA(item.Effect), item.Value)
        } else if lang == "be" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionBE(item.Effect), item.Value)
        }
        options = append(options, huh.NewOption(desc, item.ID))
    }
    
    var title string
    if lang == "en" {
        title = " Select item to use"
    } else if lang == "ua" {
        title = " Виберіть предмет для використання"
    } else if lang == "be" {
        title = " Выберыце прадмет для выкарыстання"
    }
    
    f := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[int]().
                Title(title).
                Options(options...).
                Value(&selectedID),
        ),
    )
    
    if err := f.Run(); err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    if selectedID != -1 {
        player.UseItem(selectedID)
    }
}

func getEffectDescription(effect string) string {
    switch effect {
    case "heal":
        return "Heals"
    case "damage_boost":
        return "Boosts damage by"
    case "stamina_restore":
        return "Restores stamina by"
    default:
        return effect
    }
}

func getEffectDescriptionUA(effect string) string {
    switch effect {
    case "heal":
        return "Лікує"
    case "damage_boost":
        return "Збільшує пошкодження на"
    case "stamina_restore":
        return "Відновлює витривалість на"
    default:
        return effect
    }
}

func getEffectDescriptionBE(effect string) string {
    switch effect {
    case "heal":
        return "Лякуе"
    case "damage_boost":
        return "Павялічвае пашкоджанні на"
    case "stamina_restore":
        return "Аднаўляе вынослівасць на"
    default:
        return effect
    }
}
