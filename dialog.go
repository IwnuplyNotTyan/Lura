package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
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
func dialBuffRW() {
	if lang == "en" {
		fmt.Println(style.Render("You choice random weapon. Good luck"))
	} else if lang == "ua" {
		fmt.Println(style.Render("Ви вибрали випадкову зброю. Хай вам щастить"))
	}
}

func dialBuffUW() {
	if lang == "en" {
		fmt.Println(style.Render("You add 10 damage by this upgrade, this weapon now realy strong"))
	} else if lang == "ua" {
		fmt.Println(style.Render("Ви добавили 10 пошкодження задопомогою цього покращення, ця зброя теперь насправдi сильна"))
	}
}
