package main

import (
	"fmt"
	"os"

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
	fmt.Println(style.Render("Lura ~ open source turn based rpg in CLI, fight in many locations with many monster and etc. Made with "))
}

// ASCII Art
func caveArt() {
	data, err := assetsFS.ReadFile("assets/cave.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(data))
	os.Stdout.Sync()
}

func forestArt() {
	data, err := assetsFS.ReadFile("assets/forest.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Print(string(data))
	os.Stdout.Sync()
}

// Etc
func newLine() {
	fmt.Println()
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Fight
func displayFightIntro(player *Player, monster *Monster) {
	if player.monster == false {
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
			fmt.Println(termenv.String(fmt.Sprintf("  You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Пачвар %s з'явіўся %d ХП!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
			fmt.Println(termenv.String(fmt.Sprintf("  У цябе зброя %s наносіць %d пашкоджанняй, у цябе %d ХП", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  %s з'являється з %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
			fmt.Println(termenv.String(fmt.Sprintf("  Ти володієш %s, наносиш %d пошкодження, у тебе %d здоров'я.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		}
	} else {
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
			fmt.Println(termenv.String(fmt.Sprintf("  You %s, dealing %d damage and have %d HP.", player.name, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		}
	}
}

func healDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы вылечыліся, цяпер у цябе %d ХП.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healMonsterDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр вылечыўся! Цяпер у яго %d ХП.", monster.HP)).Foreground(termenv.ANSIGreen))
	}
}

func blockDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You block the attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти блокуєш атаку!")).Foreground(termenv.ANSIYellow))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы блакуеце атаку!")).Foreground(termenv.ANSIYellow))
	}
}

func blockUDialog() {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster blocked your attack!")).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр заблокував твою атаку!")).Foreground(termenv.ANSIGreen))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр заблакаваў тваю атаку!")).Foreground(termenv.ANSIGreen))
	}
}

func blockEnemyAttack(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You blocked the enemy's attack!")).Foreground(termenv.ANSIYellow))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти заблокував атаку ворога!")).Foreground(termenv.ANSIYellow))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы заблакавалі атаку ворага!")).Foreground(termenv.ANSIYellow))
	}
}

func blockEnemyDialog(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster prepares to block!")).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр готується заблокувати!")).Foreground(termenv.ANSIGreen))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр рыхтуецца блакаваць!")).Foreground(termenv.ANSIGreen))
	}
}

func defeatMonster(monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  %s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  %s быў пераможаны\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	}
}

func staminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You have %d stamina left", player.Stamina)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  У тебе %d витривалостi залишилося", player.Stamina)).Foreground(termenv.ANSIGreen))
	} else if lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  У вас засталося %d вынослівасці", player.Stamina)).Foreground(termenv.ANSIGreen))
	}
}

func noStaminaDialog(player *Player) {
	if lang == "en" {
		fmt.Println(termenv.String("  Not enough stamina to attack!").Foreground(termenv.ANSIRed))
	} else if lang == "ua" {
		fmt.Println(termenv.String("  Недостатньо витривалості для атаки!").Foreground(termenv.ANSIRed))
	} else if lang == "be" {
		fmt.Println(termenv.String("  Недастаткова вынослівасці для атакі!").Foreground(termenv.ANSIRed))
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
	} else if lang == "be" {
		fmt.Println(termenv.String("  Бафф не быў ужыты.").Foreground(termenv.ANSIYellow))
	}
}

func currentCoins(player *Player) {
	if lang == "en" {
		fmt.Printf("  You have %d coins\n", player.Coins)
	} else if lang == "ua" {
		fmt.Printf("  У тебе %d копiйок\n", player.Coins)
	} else if lang == "be" {
		fmt.Printf("  У вас %d манет\n", player.Coins)
	}
}

func noCoinsDialog() {
	if lang == "ua" {
		fmt.Println(termenv.String("  Недостатньо копiйок.").Foreground(termenv.ANSIYellow))
	} else if lang == "en" {
		fmt.Println(termenv.String("  Not enough coins.").Foreground(termenv.ANSIYellow))
	} else if lang == "be" {
		fmt.Println(termenv.String("  Недастаткова манет.").Foreground(termenv.ANSIYellow))
	}
}
