package main

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/muesli/termenv"
)

func buffsAction(player *Player) {
	player.Coins += 10
	if lang == "en" {
		fmt.Printf("\n You have %d coins\n", player.Coins)
	} else {
		fmt.Printf("\n У тебе %d копiйок\n", player.Coins)
	}

	baff1 := getRandomBuff()
	baff2 := getRandomBuff()
	baff3 := getRandomBuff()

	var prompt promptui.Select

	if lang == "en" {
		prompt = promptui.Select{
			Label: "Select a Buff/Weapon (Upgrade)",
			Items: []string{baff1, baff2, baff3},
		}
	} else if lang == "ua" {
		prompt = promptui.Select{
			Label: "Виберіть бонус або зброю",
			Items: []string{baff1, baff2, baff3},
		}
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed:", err)
	}

	if result == "Increase HP (+2) & Reduce Damage (-1)" || result == "Додано здоров'я (+2) & Зменшено пошкодження (-1)" {
		player.HP += 2
		if player.Damage > 1 {
			player.Damage -= 1
			player.maxHP += 2
		} else {
			fmt.Println(termenv.String(" Damage cannot be reduced further!").Foreground(termenv.ANSIRed))
		}
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Increase Damage (+5) & Reduce HP (-5)" || result == "Додано пошкодження (+5) & Зменшено здоров'я (-5)" {
		player.Damage += 5
		if player.maxHP > 5 {
			player.maxHP -= 5
			player.HP -= 5
		} else {
			player.maxHP = 1
		}
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Random Weapon" || result == "Випадкова зброя" {
		weaponType, weaponDamage := getRandomWeapon()
		player.WeaponType = weaponType
		player.Damage = weaponDamage
	} else if result == "Додано витривалiсть (+10) & Зменшино пошкодження (-2)" || result == "Increase Stamina (+10) & Reduce Damage (-2)" {
		player.maxStamina += 10
	} else if result == "Add Armor (+50)" || result == "Добавити захисту (+50)" {
		player.HP += 50
	} else if result == "Upgrade Weapon" || result == "Покращити зброю" {
		if player.Coins > 20 {
			player.Damage += 10
			player.Coins -= 30
		} else if lang == "ua" {
			fmt.Println(termenv.String(" Бафф не застосовано.").Foreground(termenv.ANSIYellow))
		} else {
			fmt.Println(termenv.String(" No Buff Applied.").Foreground(termenv.ANSIYellow))
		}
	} else if lang == "ua" {
		fmt.Println(termenv.String(" Бафф не застосовано.").Foreground(termenv.ANSIYellow))
	} else {
		fmt.Println(termenv.String(" No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}
