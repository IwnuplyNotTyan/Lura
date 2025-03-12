package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(1).
	Width(35)

// Welcome
func dialWelcome() {
	fmt.Println(style.Render("Lura ~ open source turn based rpg in CLI, only you can select choose. Made with "))
}

// Buffs
func upgradeWeaponBuff(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" Weapon upgraded! Damage: %d, Coins left: %d", player.Damage, player.Coins)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Зброю покращено! Пошкодження: %d, Залишилось копiйок: %d", player.Damage, player.Coins)).Foreground(termenv.ANSIGreen))
	}
}

func increaseHPuD(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
	}
}

func increaseDuHP(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
	}
}

func increaseSuD(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, Max Stamina: %d", player.Damage, player.maxStamina)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Пошкодження: %d, Макс. витривалість: %d", player.Damage, player.maxStamina)).Foreground(termenv.ANSIGreen))
	}
}

func armorBuff(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! HP: %d", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func currentCoins(player *Player) {
	if lang == "en" {
		fmt.Printf(" You have %d coins\n", player.Coins)
	} else if lang == "ua" {
		fmt.Printf(" У тебе %d копiйок\n", player.Coins)
	}
}
