package main

import (
	"fmt"
	"time"

	"github.com/muesli/termenv"
)

var getRandomMonster string

func promptAction() string {
	result := getSelectedAttack()
	return result
}

func fight(player *Player, monster *Monster) {
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
				blockDialog()
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

		defeatMonster(monster)
		buffsAction(player)
	}
}

func fightIntro(player *Player, monster *Monster) {
	displayFightIntro(player, monster)
}

func healPlayer(player *Player) {
	player.HP = min(player.HP+15, player.maxHP)
	healDialog(player)
}

func playerAttack(player *Player, monster *Monster, playerDefending *bool, monsterDefending *bool) {
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
		blockUDialog()
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
		noStaminaDialog(player)
	}
}

func playerSkip(player *Player) {
	if player.Stamina < 100 {
		staminaDialog(player)
	}
}

func monsterTurnAction(monster *Monster, player *Player, monsterDefending *bool, playerDefending *bool, monsterAction string) {
	if monsterAction == "Defend" {
		blockEnemyDialog(monster)
		*monsterDefending = true
	} else if monsterAction == "Heal" {
		monster.HP = min(monster.HP+15, monster.maxHP)
		monster.HP = min(monster.HP+15, monster.maxHP)
		healMonsterDialog(monster)
		*monsterDefending = false
	} else {
		monsterDamage := monster.Damage + rng() + monster.LVL
		if *playerDefending {
			blockEnemyAttack(monster)
			*playerDefending = false
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
