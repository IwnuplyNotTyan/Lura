package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var (
	debugMode       = flag.Bool("debug", false, "Enable debug shell")
	specificMonster *Monster
)

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	L := lua.NewState()
	defer L.Close()
	dialWelcome()

	lang = getSelectedLanguage()
	seedData()

	registerTypes(L)
	L.SetGlobal("lang", lua.LString(lang))

	if err := AutoLoadMods(L); err != nil {
		log.Fatalf("Failed to auto-load mods: %v", err)
	}

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

	if *debugMode {
		DebugShell(L, &player)
	}

	fight(&player)
}
