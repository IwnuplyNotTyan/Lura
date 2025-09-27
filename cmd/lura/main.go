package main

import (
	"flag"
	"fmt"

	"Lura/module/data"
	"Lura/module/fight"
	"Lura/module/rng"
	"Lura/module/mods"
	"Lura/module/debug/shell"
	"Lura/module/dialog"

	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
)

var (
	debugMode       = flag.Bool("debug", false, "Enable debug shell")
	verboseMode 	= flag.Bool("v", false, "Enable verbose mode")
	specificMonster *data.Monster
)

func main() {
	flag.Parse()
	L := lua.NewState()
	defer L.Close()
	dialog.ClearScreen()

	player := data.Player{}
	cfg := data.TouchConfig(&player)

	// Исправлено: правильная работа с языком
	if cfg.Language == "" {
		fmt.Println()
		data.Lang = data.SelectLanguage()
		cfg.Language = data.Lang
		if err := data.SaveConfig(data.GetConfigPath(), cfg); err != nil {
			log.Printf("Error saving language config: %v", err)
		}
	} else {
		// Используем язык из конфига
		data.Lang = cfg.Language
	}

	data.SeedData()
	mods.RegisterTypes(L)
	L.SetGlobal("lang", lua.LString(data.Lang))

	if err := mods.AutoLoadMods(L); err != nil {
		log.Fatalf("Failed to auto-load mods: %v", err)
	}

	//if loadedMods == nil {
	//	fmt.Print(termenv.String(" ✥ ").Foreground(termenv.ANSIWhite).Bold())
	//}

	//fmt.Print(termenv.String(" 🰕 ").Foreground(termenv.ANSIMagenta).Bold())

	weaponType, weaponDamage, weaponID := rng.GetRandomWeapon()
	hp := rng.RngHp()
	st := rng.RngHp()

	player = data.Player{
		WeaponType: weaponType,
		Damage:     weaponDamage * rng.Rng(),
		HP:         hp,
		MaxHP:      hp,
		Coins:      0,
		Stamina:    st,
		MaxStamina: st,
		Heart:      1,
		Loc:        0,
		WeaponID:   weaponID,
		Monster:    false,
		Position:   0,
		Inventory: data.Inventory{
			Items:  make([]data.Item, 0),
			NextID: 1,
		},
	}

	if *debugMode {
		debug.DebugShell(L, &player, &data.Monster{})
	}

	if *verboseMode {
		data.Verbose = true
	}
	
	fight.Fight(&player, specificMonster, &cfg, &data.Weapon{})
}
