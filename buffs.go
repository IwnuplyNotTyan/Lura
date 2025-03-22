package main

import (
	"fmt"
	"math/rand"

	"github.com/muesli/termenv"
)

var (
	buff1 string
	buff2 string
	buff3 string
)

func getRandomBuff() string {
	var buffs []string

	if lang == "en" {
		buffs = []string{
			"Upgrade Weapon",
			"Random Weapon",
			"Tears",
			"Broked heart",
			"Lotus",
			//"Pearl necklace",
			"Turtle scute",
		}
	} else {
		buffs = []string{
			"Покращити зброю",
			"Випадкова зброя",
			"Розбите серце",
			"Щиток черепахи",
			"Лотос",
			//"Перлове намисто",
			"Сльози",
		}
	}
	return buffs[rand.Intn(len(buffs))]
}

func buffsAction(player *Player) {
	player.Coins += 10
	currentCoins(player)

	// Correct variable assignment
	buff1 = getRandomBuff()
	buff2 = getRandomBuff()
	buff3 = getRandomBuff()

	result := getSelectedBuff()

	if result == "Random Weapon" || result == "Випадкова зброя" {
		weaponType, weaponDamage := getRandomWeapon()
		player.WeaponType = weaponType
		player.Damage = weaponDamage
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" You found a %s! Damage: %d", weaponType, weaponDamage)).Foreground(termenv.ANSIGreen))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf(" Ти знайшов %s! Пошкодження: %d", weaponType, weaponDamage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Щиток черепахи" || result == "Turtle scute" {
		player.HP += 50
	} else if result == "Lotus" || result == "Лотус" {
		player.maxStamina += 10
	} else if result == "Tears" || result == "Сльози" {
		player.maxHP += 10
		//} else if result == "Pearl necklace" || result == "Перлове намисто" {
		//fmt.Println("placeholder")
	} else if result == "Broked heart" || result == "Розбите серце" {
		player.heart = false
	} else if result == "Upgrade Weapon" || result == "Покращити зброю" {
		player.Damage += 5
	} else {
		noBuffDialog()
	}
}
