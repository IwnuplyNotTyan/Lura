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
	} else if lang == "be" {
		buffs = []string{
			"Палепшыць зброю",
			"Выпадковая зброя",
			"Слёзы",
			"Разбітае сэрца",
			"Лотас",
			//"Перламутравае намыльнае",
			"Шчыт чарапахі",
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
	currentCoins(player)

	// Correct variable assignment
	buff1 = getRandomBuff()
	buff2 = getRandomBuff()
	buff3 = getRandomBuff()

	result := getSelectedBuff()

	if result == "Random Weapon" || result == "Випадкова зброя" || result == "Выпадковая зброя" {
		weaponType, weaponDamage := getRandomWeapon()
		player.WeaponType = weaponType
		player.Damage = weaponDamage
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" You found a %s! Damage: %d", weaponType, weaponDamage)).Foreground(termenv.ANSIGreen))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf(" Ти знайшов %s! Пошкодження: %d", weaponType, weaponDamage)).Foreground(termenv.ANSIGreen))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf(" Ты знайшоў %s! Пашкоджанні: %d", weaponType, weaponDamage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Щиток черепахи" || result == "Turtle scute" || result == "Шчыт чарапахі" {
		player.HP += 50
	} else if result == "Lotus" || result == "Лотус" || result == "Лотас" {
		player.maxStamina += 10
	} else if result == "Tears" || result == "Сльози" || result == "Слёзы" {
		player.maxHP += 10
		//} else if result == "Pearl necklace" || result == "Перлове намисто" {
		//fmt.Println("placeholder")
	} else if result == "Broked heart" || result == "Розбите серце" || result == "Разбітае сэрца" {
		player.heart = 0
	} else if result == "Upgrade Weapon" || result == "Покращити зброю" || result == "Палепшыць зброю" {
		player.Damage += 5
	} else {
		noBuffDialog()
	}
}
