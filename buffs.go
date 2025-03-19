package main

import (
	"fmt"

	"github.com/muesli/termenv"
)

var (
	buff1 string
	buff2 string
	buff3 string
)

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
	} else if result == "Черепаха щиткова" || result == "Turtle scute" {
		player.HP += 50
	} else if result == "Lotus" || result == "Лотус" {
		player.maxStamina += 10
	} else if result == "Tears" || result == "Сльози" {
		player.maxHP += 10
	} else if result == "Pearl necklace" || result == "Перлове намисто" {
		fmt.Println("placeholder")
	} else if result == "Broked heart" || result == "Розбите серце" {
		fmt.Printf("placeholder")
	} else if result == "Upgrade Weapon" || result == "Покращити зброю" {
		player.Damage += 5
	} else {
		noBuffDialog()
	}
}
