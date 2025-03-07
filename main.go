package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/muesli/termenv"
	lua "github.com/yuin/gopher-lua"
)

var term = termenv.EnvColorProfile()
var name string

func main() {
	rand.Seed(time.Now().UnixNano())
	L := lua.NewState()
	defer L.Close()
	dialWelcome()

	registerTypes(L)

	selectLanguage()
	seedData()

	L.SetGlobal("lang", lua.LString(lang))

	if err := AutoLoadMods(L); err != nil {
		log.Fatalf("Failed to auto-load mods: %v", err)
	}

	//checkAll()

	weaponType, weaponDamage := getRandomWeapon()
	player := Player{
		WeaponType: weaponType,
		Damage:     weaponDamage * rng(),
		HP:         100,
		maxHP:      100,
		Coins:      0,
		Stamina:    100,
		maxStamina: 100,
	}

	fight(&player)
}

func selectLanguage() {
	prompt := promptui.Select{
		Label: "Select a language",
		Items: []string{"English", "Українська"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	if result == "Українська" {
		lang = "ua"
	} else {
		lang = "en"
	}
}
