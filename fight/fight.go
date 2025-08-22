package fight

import (
	"fmt"
	"time"

	"Lura/data"
	"Lura/module/buffs"
	"Lura/module/rng"
	"Lura/module/dialog"

	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
)

func Fight(player *data.Player, monster *data.Monster, config *data.Config, weapon *data.Weapon) {
	//afterLoc(player)
	data.Tmp = 15
	for player.HP > 0 {
		switch player.Loc {
		case 0:
			monster = rng.GetRandomVMonster()
		case 1:
			monster = rng.GetRandomCMonster()
		case 2:
			monster = rng.GetRandomBoss()
		}
		if monster == nil {
			fmt.Println(termenv.String("No monsters found!").Foreground(termenv.ANSIYellow))
			return
		}
		monster.Position = 5

		dialog.DisplayFightIntro(player, monster)

		playerDefending := false
		monsterDefending := false
		
		for monster.HP > 0 && player.HP > 0 {
			DisplayPositions(player, monster)
			playerAction := SelectAttack(player)

			if playerAction == "Defend" || playerAction == "Захищатися" || playerAction == "Абараняцца" {
				playerDefending = true
				dialog.BlockDialog()
				if player.Position > 0 {
					player.Position--
				}
			} else if playerAction == "Heal" || playerAction == "Лікуватися" || playerAction == "Вылечвацца" {
				healPlayer(player)
				playerDefending = false
			} else if playerAction == "Attack" || playerAction == "Атакувати" || playerAction == "Атакаваць" {
				if player.Position < monster.Position-1 {
					if player.WeaponID == 5 || player.WeaponID == 6 || player.WeaponID == 10 || player.WeaponID == 8 {
						player.Position += 1
					} else {
						player.Position += 2
						if data.Lang == "en" {
							fmt.Println(termenv.String(fmt.Sprintf("󰓥  You not so close to %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
						} else if data.Lang == "ua" {
							fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти не так близько до %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
						} else if data.Lang == "be" {
							fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты не так блізка да %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
						}
					}
				}
				if player.Position == monster.Position-1 {
					playerAttack(player, monster, &monsterDefending)
				} else if player.WeaponID == 5 || player.WeaponID == 6 || player.WeaponID == 10 || player.WeaponID == 8 {
					playerAttack(player, monster, &monsterDefending)
				}
			} else if playerAction == "Skip" || playerAction == "Пропустити" || playerAction == "Прапусціць" {
				playerSkip(player)
			}

			monsterAction := enemyTurn(monster, player)
			monsterTurnAction(monster, player, &monsterDefending, &playerDefending, monsterAction)

			if player.HP <= 0 || data.Tmp == 0 {
				if player.Heart == 1 {
					if player.Score > config.Score {
						config.Score = player.Score
						if data.Lang == "en" {
							fmt.Println(termenv.String(fmt.Sprintf("  New High Score, %d", player.Score)).Foreground(termenv.ANSIBrightRed).Bold())
						} else if data.Lang == "ua" {
							fmt.Println(termenv.String(fmt.Sprintf("  Новий рекорд, %d", player.Score)).Foreground(termenv.ANSIBrightRed).Bold())
						} else if data.Lang == "be" {
							fmt.Println(termenv.String(fmt.Sprintf("  Новы рэкорд, %d", player.Score)).Foreground(termenv.ANSIBrightRed).Bold())
						}
						if err := data.SaveConfig(data.GetConfigPath(), *config); err != nil {
							log.Printf("Error saving high score: %v", err)
						}
					} else {
						fmt.Println(termenv.String(fmt.Sprintf("  %d", player.Score)).Foreground(termenv.ANSIRed).Bold())
					}
					return
				} else if player.Heart == 0 {
					player.MaxHP = player.MaxHP / 2
					player.HP = player.MaxHP
					player.Damage = player.Damage * 2
					if data.Lang == "en" {
						fmt.Println(termenv.String(fmt.Sprintf("  Your heart is broken! HP set to %d, Damage increased to %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					} else if data.Lang == "ua" {
						fmt.Println(termenv.String(fmt.Sprintf("  Ваше серце розбито! HP встановлено на %d, Пошкодження збільшено до %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					} else if data.Lang == "be" {
						fmt.Println(termenv.String(fmt.Sprintf("  Ваша сэрца разбіта! HP устаноўлена на %d, Пашкоджанні павялічаны да %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					}
					player.Heart = 1
					Fight(player, monster, config, weapon)
				}
			}
			time.Sleep(time.Second)
		}

		dialog.ClearScreen()
		player.Score += monster.Score
		player.Coins += monster.Coins
		monster.LVL += 1

		player.Position = 0
		monster.Position = 5

		dialog.DefeatMonster(monster)
		if monster.ID == 1 {
			if rng.Rng2() == 1 {
				takeWeapon(player, monster)
			}
		} else if monster.ID == 17 {
				takeWeapon(player, monster)
		}
		if player.WeaponID == 7 {
			if player.Time == 1 {
				player.Damage = monster.Damage * rng.Rng()
				player.WeaponID = 0
				player.Monster = true
				player.HP = monster.MaxHP
				player.MaxHP = monster.MaxHP
				player.Name = monster.MonsterType
				player.WeaponType = ""

				if data.Lang == "en" {
					fmt.Println(termenv.String(fmt.Sprintf("  Now you %s", player.Name)).Foreground(termenv.ANSIRed).Bold())
				} else if data.Lang == "ua" {
					fmt.Println(termenv.String(fmt.Sprintf("  Тепер ти %s", player.Name)).Foreground(termenv.ANSIRed).Bold())
				} else if data.Lang == "be" {
					fmt.Println(termenv.String(fmt.Sprintf("  Цяпер ты %s", player.Name)).Foreground(termenv.ANSIRed).Bold())
				}
			} else {
				player.Time = 1
			}
		}
		if player.HP > 0 {
		if player.Buffs == 4 {
			if !player.Monster {
				buff.BuffsAction(player)
			}
			player.Buffs = 0
			fmt.Println()
			if player.Loc == 0 {
				dialog.CaveArt()
				player.Loc = 1
			} else if player.Loc == 1 {
				player.Loc = 2
				dialog.CatArt()
			}
		} else if player.Loc == 2 {
			player.Loc = 0
			if data.Lang == "en" {
				fmt.Println(termenv.String("󰒙  Boss defeated").Foreground(termenv.ANSIBrightGreen).Bold())
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String("󰒙  Boss defeated").Foreground(termenv.ANSIBrightGreen).Bold())
			} else if data.Lang == "be" {
				fmt.Println(termenv.String("󰒙  Boss defeated").Foreground(termenv.ANSIBrightGreen).Bold())
			}
			dialog.ForestArt()
		} else {
			player.Buffs += 1
			if data.Lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Step to another location", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Крокiв до iншого мiсця", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			} else if data.Lang == "be" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Крокаў да баффу", player.Buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			}
		}}
	}

	if player.Heart == 2 {
		player.HP += 10
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

