package main

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

var getRandomMonster string

func promptAction() string {
	result := getSelectedAttack()
	// var prompt promptui.Select
	//
	//	if lang == "en" {
	//		prompt = promptui.Select{
	//			Label: "Select a card",
	//			Items: []string{"Attack", "Defend", "Heal", "Skip"},
	//		}
	//	} else {
	//
	//		prompt = promptui.Select{
	//			Label: "Вибери карту",
	//			Items: []string{"Атакувати", "Захищатися", "Лікуватися", "Пропустити"},
	//		}
	//	}
	//
	// _, result, err := prompt.Run()
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	return result
}

func fight(player *Player) {
	for player.HP > 0 {
		monster := getRandomVMonster()
		if monster == nil {
			fmt.Println(termenv.String("No monsters found!").Foreground(termenv.ANSIYellow))
			return
		}

		displayFightIntro(player, monster)

		playerDefending := false
		monsterDefending := false

		for monster.HP > 0 && player.HP > 0 {
			playerAction := promptAction()

			if playerAction == "Defend" || playerAction == "Захищатися" {
				playerDefending = true
				if lang == "en" {
					fmt.Println(termenv.String(fmt.Sprintf(" You block the attack!")).Foreground(termenv.ANSIYellow))
				} else if lang == "ua" {
					fmt.Println(termenv.String(fmt.Sprintf(" Ти блокуєш атаку!")).Foreground(termenv.ANSIYellow))
				}
			} else if playerAction == "Heal" || playerAction == "Лікуватися" {
				healPlayer(player)
				playerDefending = false
			} else if playerAction == "Attack" || playerAction == "Атакувати" {
				playerAttack(player, monster, &playerDefending, &monsterDefending)
			} else if playerAction == "Skip" || playerAction == "Пропустити" {
				playerSkip(player)
			}

			// Monster's turn
			monsterAction := enemyTurn(monster)
			monsterTurnAction(monster, player, &monsterDefending, &playerDefending, monsterAction)

			// Check if player died
			if player.HP <= 0 {
				fmt.Println(termenv.String(" ").Foreground(termenv.ANSIBrightRed).Bold())
				return
			}

			time.Sleep(time.Second)
		}

		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("%s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		}
		buffsAction(player)
	}
}

func displayFightIntro(player *Player, monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf(" You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Дикий %s з'являється з %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf(" Ти володієш %s, наносиш %d пошкодження, у тебе %d здоров'я.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func healPlayer(player *Player) {
	player.HP = min(player.HP+15, player.maxHP)
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func playerAttack(player *Player, monster *Monster, playerDefending *bool, monsterDefending *bool) {
	// Find the equipped weapon
	var weapon *Weapon
	for _, w := range weapons {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	if weapon == nil {
		weapon = &Weapon{WeaponType: "Fists", Damage: 2, Stamina: 0}
	}

	playerDamage := player.Damage + rng()
	if *monsterDefending {
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" The monster blocked your attack!")).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Монстр заблокував твою атаку!")).Foreground(termenv.ANSIGreen))
		}
		*monsterDefending = false
	} else if player.Stamina >= weapon.Stamina {
		player.Stamina -= weapon.Stamina
		monster.HP -= playerDamage
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥 You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥 Ти атакував %s з силою %d! Тепер в нього %d здоров'я. У тебе залишилось %d витривалостi", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		}
	} else {
		if lang == "en" {
			fmt.Println(termenv.String(" Not enough stamina to attack!").Foreground(termenv.ANSIRed))
		} else {
			fmt.Println(termenv.String(" Недостатньо витривалості для атаки!").Foreground(termenv.ANSIRed))
		}
	}
}

func playerSkip(player *Player) {
	if player.Stamina < 100 {
		player.Stamina = min(player.Stamina+20, player.maxStamina)
	}
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You have %d stamina left", player.Stamina)).Foreground(termenv.ANSIGreen))
	} else if lang == "ua" {
		fmt.Println(termenv.String(fmt.Sprintf(" У тебе %d витривалостi залишилося", player.Stamina)).Foreground(termenv.ANSIGreen))
	}
}

func monsterTurnAction(monster *Monster, player *Player, monsterDefending *bool, playerDefending *bool, monsterAction string) {
	if monsterAction == "Defend" {
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" The monster prepares to block!")).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Монстр готується заблокувати!")).Foreground(termenv.ANSIGreen))
		}
		*monsterDefending = true
	} else if monsterAction == "Heal" {
		monster.HP = min(monster.HP+15, monster.maxHP)
		monster.HP = min(monster.HP+15, monster.maxHP)
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
		}
		*monsterDefending = false
	} else {
		monsterDamage := monster.Damage + rng() + monster.LVL
		if *playerDefending {
			if lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf(" You blocked the enemy's attack!")).Foreground(termenv.ANSIYellow))
			} else if lang == "ua" {
				fmt.Println(termenv.String(fmt.Sprintf(" Ти заблокував атаку ворога!")).Foreground(termenv.ANSIYellow))
			}
			*playerDefending = false // Reset defense after blocking
		} else {
			player.HP -= monsterDamage
			if lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 The %s attacks you for %d damage! You now have %d HP.", monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
			} else {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.", monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
			}
		}
	}
}

func enemyTurn(monster *Monster) string {
	rngChoice := rng() % 3

	switch rngChoice {
	case 0:
		return "Attack"
	case 1:
		return "Defend"
	case 2:
		return "Heal"
	default:
		return "Attack"
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
