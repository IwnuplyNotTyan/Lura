package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(1).
	Height(4).
	Width(40)

func getLine(lines []string, index int) string {
	if index < len(lines) {
		return lines[index]
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Welcome
func dialWelcome() {
	ascii := `[38;2;26;21;28m:[0m[38;2;48;52;99m:[0m[38;2;85;101;207m:[0m[38;2;85;101;207m:[0m[38;2;48;52;99m:[0m[38;2;26;21;28m:[0m[38;2;30;26;40m:[0m[38;2;35;34;56m:[0m[38;2;34;32;52m:[0m[38;2;34;32;52m:[0m
[38;2;69;92;248m:[0m[38;2;134;145;177m:[0m[38;2;232;226;65m:[0m[38;2;238;233;81m:[0m[38;2;140;154;195m:[0m[38;2;56;75;211m:[0m[38;2;35;39;99m:[0m[38;2;32;27;28m:[0m[38;2;31;27;40m:[0m[38;2;35;34;56m:[0m
[38;2;26;21;28m:[0m[38;2;48;52;99m:[0m[38;2;85;101;207m:[0m[38;2;85;102;207m:[0m[38;2;51;55;103m:[0m[38;2;22;16;17m:[0m[38;2;24;18;21m:[0m[38;2;47;50;93m:[0m[38;2;79;93;189m:[0m[38;2;97;118;243m:[0m
[38;2;97;118;243m:[0m[38;2;79;93;189m:[0m[38;2;47;50;93m:[0m[38;2;24;19;21m:[0m[38;2;28;22;17m:[0m[38;2;36;42;103m:[0m[38;2;62;79;207m:[0m[38;2;130;145;207m:[0m[38;2;221;219;100m:[0m[38;2;255;253;30m:[0m
[38;2;255;253;30m:[0m[38;2;221;219;100m:[0m[38;2;130;145;207m:[0m[38;2;62;79;207m:[0m[38;2;36;42;103m:[0m[38;2;28;22;17m:[0m[38;2;24;19;21m:[0m[38;2;47;50;93m:[0m[38;2;79;93;189m:[0m[38;2;97;118;243m:[0m
[38;2;97;120;248m:[0m[38;2;75;88;177m:[0m[38;2;38;36;64m:[0m[38;2;43;44;80m:[0m[38;2;81;97;195m:[0m[38;2;86;104;211m:[0m[38;2;48;52;99m:[0m[38;2;26;21;28m:[0m[38;2;30;26;40m:[0m[38;2;35;34;56m:[0m
[38;2;32;27;28m:[0m[38;2;35;39;99m:[0m[38;2;56;75;211m:[0m[38;2;140;154;195m:[0m[38;2;236;231;77m:[0m[38;2;236;231;77m:[0m[38;2;140;154;195m:[0m[38;2;56;75;211m:[0m[38;2;35;39;99m:[0m[38;2;32;27;28m:[0m
`
	text := style.Render("Lura ~ open source turn based rpg in CLI, only you can select choose. Made with ")

	styledText := style.Render(text)

	linesLeft := strings.Split(styledText, "\n")
	linesRight := strings.Split(ascii, "\n")

	var output strings.Builder
	maxLines := max(len(linesLeft), len(linesRight))

	for i := 0; i < maxLines; i++ {
		left := getLine(linesLeft, i)
		right := getLine(linesRight, i)

		// Only print if at least one side has content
		if strings.TrimSpace(left) != "" || strings.TrimSpace(right) != "" {
			output.WriteString(fmt.Sprintf("%-40s %s\n", left, right))
		}
	}

	fmt.Println(output.String())
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
		fmt.Println(termenv.String(fmt.Sprintf("  A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf("  You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Дикий %s з'являється з %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf("  Ти володієш %s, наносиш %d пошкодження, у тебе %d здоров'я.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healMonsterDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
	}
}

func blockDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You block the attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти блокуєш атаку!")).Foreground(termenv.ANSIYellow))
	}
}

func blockUDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster blocked your attack!")).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр заблокував твою атаку!")).Foreground(termenv.ANSIGreen))
	}
}

func blockEnemyAttack(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You blocked the enemy's attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти заблокував атаку ворога!")).Foreground(termenv.ANSIYellow))
	}
}

func blockEnemyDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster prepares to block!")).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр готується заблокувати!")).Foreground(termenv.ANSIGreen))
	}
}

func defeatMonster(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  %s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	}
}

func staminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You have %d stamina left", player.Stamina)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  У тебе %d витривалостi залишилося", player.Stamina)).Foreground(termenv.ANSIGreen))
	}
}

func noStaminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String("  Not enough stamina to attack!").Foreground(termenv.ANSIRed))
	} else {
		fmt.Println(termenv.String("  Недостатньо витривалості для атаки!").Foreground(termenv.ANSIRed))
	}
}

// Buffs
func armorBuff(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  Buff Applied! HP: %d", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("  Бафф застосовано! Здоров'я: %d", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func noBuffDialog() {
	if lang == "ua" {
		fmt.Println(termenv.String("  Бафф не застосовано.").Foreground(termenv.ANSIYellow))
	} else if lang == "en" {
		fmt.Println(termenv.String("  No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func currentCoins(player *Player) {
	if lang == "en" {
		fmt.Printf("  You have %d coins\n", player.Coins)
	} else if lang == "ua" {
		fmt.Printf("  У тебе %d копiйок\n", player.Coins)
	}
}

func noCoinsDialog() {
	if lang == "ua" {
		fmt.Println(termenv.String("  Недостатньо копiйок.").Foreground(termenv.ANSIYellow))
	} else if lang == "en" {
		fmt.Println(termenv.String("  Not enough coins.").Foreground(termenv.ANSIYellow))
	}
}
