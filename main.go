package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/muesli/termenv"
)

var term = termenv.EnvColorProfile()

func main() {
	rand.Seed(time.Now().UnixNano())

	selectLanguage()
	seedData()

	weaponType, weaponDamage := getRandomWeapon()
	player := Player{WeaponType: weaponType, Damage: weaponDamage * rng(), HP: 100, maxHP: 100, Coins: 0, Stamina: 100, maxStamina: 100}

	fight(&player)
}

func selectLanguage() {
	prompt := promptui.Select{
		Label: "Select a language",
		Items: []string{"English", "Українська"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	if result == "Українська" {
		lang = "ua"
	} else if result == "English" {
		lang = "en"
	}
}
