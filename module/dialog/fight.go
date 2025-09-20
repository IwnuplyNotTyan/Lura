package dialog

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"Lura/data"
	"Lura/assets"

	"github.com/muesli/termenv"
)

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

func MissDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("󰓥  Мимо").Foreground(termenv.ANSIBlue))
		case "ua":
			fmt.Println(termenv.String("󰓥  Мимо").Foreground(termenv.ANSIBlue))
		case "be":
			fmt.Println(termenv.String("󰓥  Cпадарыня").Foreground(termenv.ANSIBlue))
		default:
			fmt.Println(termenv.String("󰓥  Miss").Foreground(termenv.ANSIBlue))
	}
}

func HealDialog(player *data.Player) {
	fmt.Println(termenv.String(fmt.Sprintf("  %d / %d", player.HP, player.MaxHP)).Foreground(termenv.ANSIGreen))
}

func HealMonsterDialog(monster *data.Monster) {
	fmt.Println(termenv.String(fmt.Sprintf("  %d / %d", monster.HP, monster.MaxHP)).Foreground(termenv.ANSIGreen))
}

func BlockDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Вы блокируете атаку!").Foreground(termenv.ANSIYellow))
		case "ua":
			fmt.Println(termenv.String("  Ти блокуєш атаку!").Foreground(termenv.ANSIYellow))
		case "be":
			fmt.Println(termenv.String("  Вы блакуеце атаку!").Foreground(termenv.ANSIYellow))
		default:
			fmt.Println(termenv.String("  You block the attack!").Foreground(termenv.ANSIYellow))
	}
}

func BlockUDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Монстр заблокировал вашу атаку!").Foreground(termenv.ANSIGreen))
		case "ua":
			fmt.Println(termenv.String("  Монстр заблокував твою атаку!").Foreground(termenv.ANSIGreen))
		case "be":
			fmt.Println(termenv.String("  Монстр заблакаваў тваю атаку!").Foreground(termenv.ANSIGreen))
		default:
			fmt.Println(termenv.String("  The monster blocked your attack!").Foreground(termenv.ANSIGreen))
	}
}

func BlockEnemyAttack() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Вы блокируете атаку врага!").Foreground(termenv.ANSIYellow))
		case "ua":
			fmt.Println(termenv.String("  Ти блокуєш атаку ворога!").Foreground(termenv.ANSIYellow))
		case "be":
			fmt.Println(termenv.String("  Вы блакуеце атаку ворага!").Foreground(termenv.ANSIYellow))
		default:
			fmt.Println(termenv.String("  You block the enemy's attack!").Foreground(termenv.ANSIYellow))
	}
}

func BlockEnemyDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Монстр готовится блокировать!").Foreground(termenv.ANSIGreen))
		case "ua":
			fmt.Println(termenv.String("  Монстр готується заблокувати!").Foreground(termenv.ANSIGreen))
		case "be":
			fmt.Println(termenv.String("  Монстр рыхтуецца блакаваць!").Foreground(termenv.ANSIGreen))
		default:
			fmt.Println(termenv.String("  The monster prepares to block!").Foreground(termenv.ANSIGreen))
	}
}

func DefeatMonster(monster *data.Monster) {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String(fmt.Sprintf("  %s был побежден!", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		case "ua":
			fmt.Println(termenv.String(fmt.Sprintf("  %s был побежден!", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		case "be":
			fmt.Println(termenv.String(fmt.Sprintf("  %s быў пераможаны!", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		default:
			fmt.Println(termenv.String(fmt.Sprintf("  The %s has been defeated!", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
	}
}

func StaminaDialog(player *data.Player) {
	if !player.Monster {
			fmt.Println(termenv.String(fmt.Sprintf("  %d / %d", player.Stamina, player.MaxStamina)).Foreground(termenv.ANSIGreen))
	}
}

func NoStaminaDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Недостаточно выносливости").Foreground(termenv.ANSIRed))
		case "ua":
			fmt.Println(termenv.String("  Недостатньо витривалості").Foreground(termenv.ANSIRed))
		case "be":
			fmt.Println(termenv.String("  Недастаткова вынослівасці").Foreground(termenv.ANSIRed))
		default:
			fmt.Println(termenv.String("  Not enough stamina").Foreground(termenv.ANSIRed))
	}
}

func GatoDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String(fmt.Sprintf("  %d шагов до смерти", data.Tmp)).Foreground(termenv.ANSIBlue))
		case "ua":
			fmt.Println(termenv.String(fmt.Sprintf("  %d крокiв до смерті", data.Tmp)).Foreground(termenv.ANSIBlue))
		case "be":
			fmt.Println(termenv.String(fmt.Sprintf("  %d крокаў да мертвості", data.Tmp)).Foreground(termenv.ANSIBlue))
		default:
			fmt.Println(termenv.String(fmt.Sprintf("  %d steps to die", data.Tmp)).Foreground(termenv.ANSIBlue))

	}
}

func FarDialog(monster *data.Monster) {
	    switch data.Lang {
		case "ru":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s слишком далеко для атаки!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
		case "ua":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадто далеко для атаки!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
		case "be":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадта далёка для атакі!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
		default:
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s is too far away to attack!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
	    }
}

func PlayerAttackDialog(monster *data.Monster, player *data.Player, playerDamage int) {
		switch data.Lang {
			case "ru":
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атаковал %s на %d урона! У него %d здоровья. Осталось %d выносливости.", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
			case "be":
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атакаваў %s на %d здароўя! Цяпер у яго %d ХП. Засталось %d вынослівасьці.", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
			case "ua": 
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти атакував %s з силою %d! Тепер в нього %d здоров'я. У тебе залишилось %d витривалостi", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
			default:
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		}

}

func TooFarDialog(player *data.Player) {
		switch data.Lang {
			case "ru":
				fmt.Println(termenv.String("󰓥  Ты слишком далеко от монстра чтобы атаковать своим оружием!").Foreground(termenv.ANSIYellow))
			case "ua":
				fmt.Println(termenv.String("󰓥  Ти занадто далеко від монстра щоб атакувати своїм оружием!").Foreground(termenv.ANSIYellow))
			case "be":
				fmt.Println(termenv.String("󰓥  Ты занадта далёка ад монстра каб атакаваць сваім оружием!").Foreground(termenv.ANSIYellow))
			default:
				fmt.Println(termenv.String("󰓥  You are too far away from the monster to attack with your weapon!").Foreground(termenv.ANSIYellow))
		}
}

func NotCloseDialog(monster *data.Monster) {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты не так близко к %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
		case "ua":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти не так близько до %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
		case "be":
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты не так блізка да %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
		default:
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You not so close to %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
	}
}

func AttackDialog(monster *data.Monster, player *data.Player, monsterDamage int) {
		switch data.Lang {
			case "ru":
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атаковал %s на %d урона! У тебе %d здоровья.",
					monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
			case "ua":
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.",
					monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
			case "be":
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s атакаваў цябе на %d здароўя! Цяпер у цябе %d ХП.",
					monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
			default:
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s attacks you for %d damage! You now have %d HP.",
					monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
		}
}

func BossDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("󰒙  Босс побежден").Foreground(termenv.ANSIBrightGreen).Bold())
		case "ua":
			fmt.Println(termenv.String("󰒙  Босс переможений").Foreground(termenv.ANSIBrightGreen).Bold())
		case "be":
			fmt.Println(termenv.String("󰒙  Босс пераможаны").Foreground(termenv.ANSIBrightGreen).Bold())
		default:
			fmt.Println(termenv.String("󰒙  Boss defeated").Foreground(termenv.ANSIBrightGreen).Bold())
	}
}

func BuffStepsDialog(player *data.Player) {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String(fmt.Sprintf("  %d шагов до другого локации", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
		case "ua":
			fmt.Println(termenv.String(fmt.Sprintf("  %d крокiв до iншого локацiї", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
		case "be":
			fmt.Println(termenv.String(fmt.Sprintf("  %d крокаў да iншага локацы", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
		default:
			fmt.Println(termenv.String(fmt.Sprintf("  %d Step to another location", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
	}
}
