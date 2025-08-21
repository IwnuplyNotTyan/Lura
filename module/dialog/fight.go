package dialog

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"Lura/data"
	"Lura/assets"

	"github.com/muesli/termenv"
)

// ASCII Art

// Fight
func DisplayFightIntro(player *data.Player, monster *data.Monster) {
	var text string
	if !player.Monster {
		text = fmt.Sprintf(" : %d  : %d 󰓥 : %d 󱡅 : %s\n : %d 󰓥 : %d 󰙊 : %s", player.HP, player.Stamina, player.Damage, player.WeaponType, monster.HP, monster.Damage, monster.MonsterType)
	} else {
		text = fmt.Sprintf(" : %d  : %d 󰓥 : %d 󰙊 : %s\n : %d 󰓥 : %d 󰙊 : %s", player.HP, player.Stamina, player.Damage, player.Name, monster.HP, monster.Damage, monster.MonsterType)
	}
	maxLength := 68
	truncatedText := text
	
	if utf8.RuneCountInString(text) > maxLength {
		if maxLength <= 3 {
			truncatedText = "..."
		} else {
			runes := []rune(text)
			truncatedText = string(runes[:maxLength-3]) + "..."
		}
	}
	
	lipText := statStyle.Render(truncatedText)
	
	filename := fmt.Sprintf("preview/monster/%d.txt", monster.ID)
	content, _ := asset.FS.ReadFile(filename)
	
	linesLeft := strings.Split(lipText, "\n")
	linesRight := strings.Split(string(content), "\n")
	
	var output strings.Builder
	maxLines := max(len(linesLeft), len(linesRight))
	
	for i := 0; i < maxLines; i++ {
		left := GetLine(linesLeft, i)
		right := GetLine(linesRight, i)
		output.WriteString(fmt.Sprintf("%-40s %s\n", left, right))
	}
	
	fmt.Print(output.String())
}

func HealDialog(player *data.Player) {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы вылечыліся, цяпер у цябе %d ХП.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func HealMonsterDialog(monster *data.Monster) {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр вылечыўся! Цяпер у яго %d ХП.", monster.HP)).Foreground(termenv.ANSIGreen))
	}
}

func BlockDialog() {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You block the attack!")).Foreground(termenv.ANSIYellow))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти блокуєш атаку!")).Foreground(termenv.ANSIYellow))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы блакуеце атаку!")).Foreground(termenv.ANSIYellow))
	}
}

func BlockUDialog() {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster blocked your attack!")).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр заблокував твою атаку!")).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр заблакаваў тваю атаку!")).Foreground(termenv.ANSIGreen))
	}
}

func BlockEnemyAttack() {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  You blocked the enemy's attack!")).Foreground(termenv.ANSIYellow))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Ти заблокував атаку ворога!")).Foreground(termenv.ANSIYellow))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Вы заблакавалі атаку ворага!")).Foreground(termenv.ANSIYellow))
	}
}

func BlockEnemyDialog() {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The monster prepares to block!")).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр готується заблокувати!")).Foreground(termenv.ANSIGreen))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  Монстр рыхтуецца блакаваць!")).Foreground(termenv.ANSIGreen))
	}
}

func DefeatMonster(monster *data.Monster) {
	if data.Lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("  The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf("  %s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	} else if data.Lang == "be" {
		fmt.Println(termenv.String(fmt.Sprintf("  %s быў пераможаны\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	}
}

func StaminaDialog(player *data.Player) {
	if !player.Monster {
		if data.Lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  You have %d stamina left", player.Stamina)).Foreground(termenv.ANSIGreen))
		} else if data.Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  У тебе %d витривалостi залишилося", player.Stamina)).Foreground(termenv.ANSIGreen))
		} else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  У вас засталося %d вынослівасці", player.Stamina)).Foreground(termenv.ANSIGreen))
		}
	}
}

func NoStaminaDialog() {
	if data.Lang == "en" {
		fmt.Println(termenv.String("  Not enough stamina to attack!").Foreground(termenv.ANSIRed))
	} else if data.Lang == "ua" {
		fmt.Println(termenv.String("  Недостатньо витривалості для атаки!").Foreground(termenv.ANSIRed))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String("  Недастаткова вынослівасці для атакі!").Foreground(termenv.ANSIRed))
	}
}
