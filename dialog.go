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

// Selectors
func languageModel() model {
	return model{
		choices: []string{"English", "Українська"},
	}
}

func attackModel() model {
	if lang == "en" {
		return model{
			choices: []string{"Attack", "Defend", "Heal", "Skip"},
		}
	} else {
		return model{
			choices: []string{"Атакувати", "Захищатися", "Лікуватися", "Пропустити"},
		}
	}
}

func buffsModel() model {
	return model{
		choices: []string{buff1, buff2, buff3},
	}
}

// Fight
func displayFightIntro(player *Player, monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf(" You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Дикий %s з'являється з %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf(" Ти володієш %s, наносиш %d пошкодження, у тебе %d здоров'я.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healMonsterDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
	}
}

func blockDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You block the attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf(" Ти блокуєш атаку!")).Foreground(termenv.ANSIYellow))
	}
}

func blockUDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" The monster blocked your attack!")).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Монстр заблокував твою атаку!")).Foreground(termenv.ANSIGreen))
	}
}

func blockEnemyAttack(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You blocked the enemy's attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf(" Ти заблокував атаку ворога!")).Foreground(termenv.ANSIYellow))
	}
}

func blockEnemyDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" The monster prepares to block!")).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Монстр готується заблокувати!")).Foreground(termenv.ANSIGreen))
	}
}

func defeatMonster(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("%s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	}
}

func staminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You have %d stamina left", player.Stamina)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf(" У тебе %d витривалостi залишилося", player.Stamina)).Foreground(termenv.ANSIGreen))
	}
}

func noStaminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(" Not enough stamina to attack!").Foreground(termenv.ANSIRed))
	} else {
		fmt.Println(termenv.String(" Недостатньо витривалості для атаки!").Foreground(termenv.ANSIRed))
	}
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

func noBuffDialog() {
	if lang == "ua" {
		fmt.Println(termenv.String(" Бафф не застосовано.").Foreground(termenv.ANSIYellow))
	} else if lang == "en" {
		fmt.Println(termenv.String(" No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func currentCoins(player *Player) {
	if lang == "en" {
		fmt.Printf(" You have %d coins\n", player.Coins)
	} else if lang == "ua" {
		fmt.Printf(" У тебе %d копiйок\n", player.Coins)
	}
}

func noCoinsDialog() {
	if lang == "ua" {
		fmt.Println(termenv.String(" Недостатньо копiйок.").Foreground(termenv.ANSIYellow))
	} else if lang == "en" {
		fmt.Println(termenv.String(" Not enough coins.").Foreground(termenv.ANSIYellow))
	}
}
