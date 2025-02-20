package main

import (
	"fmt"

	"github.com/muesli/termenv"
)

func checkAllWeapons() {
	if lang == "en" {
		fmt.Println(termenv.String("\n All Weapons:").Foreground(termenv.ANSIBlue).Bold())
	} else {
		fmt.Println(termenv.String("\n Вся зброя:").Foreground(termenv.ANSIBlue).Bold())
	}

	for _, weapon := range weapons {
		if lang == "en" {
			fmt.Printf(" Weapon: %s, Damage: %d, Stamina Cost: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
		} else {
			fmt.Printf(" Зброя: %s, Пошкодження: %d, Витрати витривалості: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
		}
	}
}

func checkAllMonsters() {
	if lang == "en" {
		fmt.Println(termenv.String("\n All Monsters:").Foreground(termenv.ANSIBlue).Bold())
	} else {
		fmt.Println(termenv.String("\n Всі монстри:").Foreground(termenv.ANSIBlue).Bold())
	}

	for _, monster := range monsters {
		if lang == "en" {
			fmt.Printf(" Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
		} else {
			fmt.Printf(" Монстр: %s, Здоров'я: %d, Пошкодження: %d, Рівень: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
		}
	}
}

func checkAll() {
	checkAllWeapons()
	checkAllMonsters()
}
