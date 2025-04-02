package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	fmt.Printf("\n")

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
		monster:    false,
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

func InventoryMenu(player *Player) {
	for {
		var choice string
		options := []huh.Option[string]{}

		if lang == "en" {
			options = append(options,
				huh.NewOption("View inventory", "view"),
				huh.NewOption("Use items", "use"),
				huh.NewOption("Drop items", "drop"),
				huh.NewOption("Sell items", "sell"),
				huh.NewOption("Exit", "exit"),
			)
		} else if lang == "ua" {
			options = append(options,
				huh.NewOption("Переглянути інвентар", "view"),
				huh.NewOption("Використати предмети", "use"),
				huh.NewOption("Видалити предмети", "drop"),
				huh.NewOption("Продати предмети", "sell"),
				huh.NewOption("Вийти", "exit"),
			)
		} else if lang == "be" {
			options = append(options,
				huh.NewOption("Прагледзець інвентар", "view"),
				huh.NewOption("Выкарыстаць прадметы", "use"),
				huh.NewOption("Выдаліць прадметы", "drop"),
				huh.NewOption("Прадаць прадметы", "sell"),
				huh.NewOption("Выйсці", "exit"),
			)
		}

		var title string
		if lang == "en" {
			title = " Inventory Menu"
		} else if lang == "ua" {
			title = " Меню інвентарю"
		} else if lang == "be" {
			title = " Меню інвентара"
		}

		err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(title).
					Options(options...).
					Value(&choice),
			),
		).Run()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch choice {
		case "view":
			ShowInventory(player)
		case "use":
			UseMultipleItems(player)
		case "drop":
			DropItems(player)
		case "sell":
			SellItems(player)
		case "exit":
			return
		}
	}
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

func MultiSelectItems(player *Player, action string) []int {
	if len(player.Inventory.Items) == 0 {
		if lang == "en" {
			fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
		} else if lang == "ua" {
			fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
		} else if lang == "be" {
			fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
		}
		return nil
	}

	var selectedIDs []int
	options := make([]huh.Option[int], len(player.Inventory.Items))

	// Create options for each item
	for i, item := range player.Inventory.Items {
		var desc string
		if lang == "en" {
			desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescription(item.Effect), item.Value)
		} else if lang == "ua" {
			desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionUA(item.Effect), item.Value)
		} else if lang == "be" {
			desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionBE(item.Effect), item.Value)
		}
		options[i] = huh.NewOption(desc, item.ID)
	}

	// Determine title based on action and language
	var title string
	switch action {
	case "use":
		if lang == "en" {
			title = " Select items to use"
		} else if lang == "ua" {
			title = " Виберіть предмети для використання"
		} else if lang == "be" {
			title = " Выберыце прадметы для выкарыстання"
		}
	case "drop":
		if lang == "en" {
			title = " Select items to drop"
		} else if lang == "ua" {
			title = " Виберіть предмети для видалення"
		} else if lang == "be" {
			title = " Выберыце прадметы для выдалення"
		}
	case "sell":
		if lang == "en" {
			title = " Select items to sell"
		} else if lang == "ua" {
			title = " Виберіть предмети для продажу"
		} else if lang == "be" {
			title = " Выберыце прадметы для продажу"
		}
	default:
		if lang == "en" {
			title = " Select items"
		} else if lang == "ua" {
			title = " Виберіть предмети"
		} else if lang == "be" {
			title = " Выберыце прадметы"
		}
	}

	// Create the form with multi-select
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[int]().
				Title(title).
				Options(options...).
				Value(&selectedIDs),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return selectedIDs
}

// Example usage functions
func UseMultipleItems(player *Player) {
	selectedIDs := MultiSelectItems(player, "use")
	if len(selectedIDs) == 0 {
		return
	}

	for _, id := range selectedIDs {
		player.UseItem(id)
	}
}

func DropItems(player *Player) {
	selectedIDs := MultiSelectItems(player, "drop")
	if len(selectedIDs) == 0 {
		return
	}

	var confirm bool
	var confirmTitle string
	if lang == "en" {
		confirmTitle = fmt.Sprintf("Are you sure you want to drop %d items?", len(selectedIDs))
	} else if lang == "ua" {
		confirmTitle = fmt.Sprintf("Ви впевнені, що хочете видалити %d предметів?", len(selectedIDs))
	} else if lang == "be" {
		confirmTitle = fmt.Sprintf("Вы ўпэўнены, што хочаце выдаліць %d прадметаў?", len(selectedIDs))
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(confirmTitle).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirm {
		for _, id := range selectedIDs {
			player.Inventory.RemoveItem(id, 1)
		}

		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  Dropped %d items", len(selectedIDs))).Foreground(termenv.ANSIGreen))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  Видалено %d предметів", len(selectedIDs))).Foreground(termenv.ANSIGreen))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Выдалена %d прадметаў", len(selectedIDs))).Foreground(termenv.ANSIGreen))
		}
	}
}

func SellItems(player *Player) {
	selectedIDs := MultiSelectItems(player, "sell")
	if len(selectedIDs) == 0 {
		return
	}

	// Calculate total value
	totalValue := 0
	for _, id := range selectedIDs {
		for _, item := range player.Inventory.Items {
			if item.ID == id {
				totalValue += item.Price / 2 // Sell for half price
				break
			}
		}
	}

	var confirm bool
	var confirmTitle string
	if lang == "en" {
		confirmTitle = fmt.Sprintf("Sell %d items for %d coins?", len(selectedIDs), totalValue)
	} else if lang == "ua" {
		confirmTitle = fmt.Sprintf("Продати %d предметів за %d копійок?", len(selectedIDs), totalValue)
	} else if lang == "be" {
		confirmTitle = fmt.Sprintf("Прадаць %d прадметаў за %d манет?", len(selectedIDs), totalValue)
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(confirmTitle).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm),
		),
	).Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirm {
		for _, id := range selectedIDs {
			player.Inventory.RemoveItem(id, 1)
		}
		player.Coins += totalValue

		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  Sold items for %d coins", totalValue)).Foreground(termenv.ANSIGreen))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  Предметів продано за %d копійок", totalValue)).Foreground(termenv.ANSIGreen))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Прадметаў прададзена за %d манет", totalValue)).Foreground(termenv.ANSIGreen))
		}
	}
}
