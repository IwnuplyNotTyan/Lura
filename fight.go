package main

import (
	"fmt"
	"strings"
	"time"

	"Lura/data"
	"Lura/module/buffs"
	"Lura/module/rng"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
)

var (
	getRandomMonster string
	monster          string
	Attack           string
	Heal             string
	Defend           string
	Skip             string
)

func displayPositions(player *data.Player, monster *data.Monster) {
    positions := make([]string, 6)
    for i := range positions {
        positions[i] = " "
    }
    
    positions[player.Position] = " "
    positions[monster.Position] = " "
    
    fmt.Println(strings.Join(positions, ""))
}

func takeWeapon(player *data.Player, monster *data.Monster) {
	var confirm bool
	var a, b, c string

	switch {
	case data.Lang == "ua":
		a = "Ви хочете взяти зброю?"
		b = "Так"
		c = "Ні"
	case data.Lang == "be":
		a = "Вы хочаце ўзяць зброю?"
		b = "Так"
		c = "Не"
	default:
		a = "Do you want to take the weapon?"
		b = "Yes"
		c = "No"
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(a).
				Affirmative(b).
				Negative(c).
				Value(&confirm),
		),
	).Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirm {
		if monster.ID == 17 {
			rng.GetLanter(player)
		} else if monster.ID == 1 {
			rng.GetMusket(player)
		}
	}
}

func afterLoc(player *data.Player) {
var confirm string

err := huh.NewSelect[string]().
	Title("Good night! What want?").
	Options(
		huh.NewOption("Sleep", "sleep"),
		huh.NewOption("Cooking", "cook"),
		huh.NewOption("Crafting", "craft"),
		huh.NewOption("Mining", "mine"),
	).
	Value(&confirm).
	Run()

if err != nil {
	log.Info(err)
	return
}

if confirm == "sleep" {
	log.Info(player.Stamina)
	staminaSleep := player.MaxStamina + 20
	player.Stamina = staminaSleep
	log.Info(player.Stamina)
} else if confirm == "craft" {
	crafting(player)
} else if confirm == "mine" {
	log.Info("Mining...")
} else if confirm == "cook" {
	log.Info("Cooking...")
} else {
	log.Info("Invalid selection")
}
}

func crafting(player *data.Player) {
    var selections []string 
    err := huh.NewMultiSelect[string]().
        Title("Crafting").
        Options(
            huh.NewOption("Bell", "bell"),
        ).
        Value(&selections).
        Run()
    if err != nil {
        log.Info(err)
        return
    }
    
    for _, selection := range selections {
        if selection == "bell" {
            log.Info("Crafting a bell")
        }
    }
}

func selectAttack(player *data.Player) string {
	var selectedAttack string
	if data.Lang == "ua" {
		Attack = "Атакувати"
		Heal = "Лікуватися"
		Defend = "Захищатися"
		Skip = "Пропустити"
	} else if data.Lang == "be" {
		Attack = "Атакаваць"
		Heal = "Вылечвацца"
		Defend = "Абараняцца"
		Skip = "Прапусціць"
	} else {
		Attack = "Attack"
		Defend = "Defend"
		Heal = "Heal"
		Skip = "Skip"
	}
	
	var f *huh.Form
	if !player.Monster {
		f = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(" Select action").
					Options(
						huh.NewOption(Attack, Attack),
						huh.NewOption(Defend, Defend),
						huh.NewOption(Heal, Heal),
						huh.NewOption(Skip, Skip),
					).
					Value(&selectedAttack),
			),
		)
	} else {
		f = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(" Select action").
					Options(
						huh.NewOption(Attack, Attack),
						huh.NewOption(Heal, Heal),
						huh.NewOption(Defend, Defend),
					).
					Value(&selectedAttack),
			),
		)
	}
	
	if err := f.Run(); err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	clearScreen()
	return selectedAttack
}

func fight(player *data.Player, monster *data.Monster, config *data.Config, weapon *data.Weapon) {
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

		displayFightIntro(player, monster)

		playerDefending := false
		monsterDefending := false
		
		for monster.HP > 0 && player.HP > 0 {
			displayPositions(player, monster)
			playerAction := selectAttack(player)

			if playerAction == "Defend" || playerAction == "Захищатися" || playerAction == "Абараняцца" {
				playerDefending = true
				blockDialog()
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
					fight(player, monster, config, weapon)
				}
			}
			time.Sleep(time.Second)
		}

		clearScreen()
		player.Score += monster.Score
		player.Coins += monster.Coins
		monster.LVL += 1

		player.Position = 0
		monster.Position = 5

		defeatMonster(monster)
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
				caveArt()
				player.Loc = 1
			} else if player.Loc == 1 {
				player.Loc = 2
				catArt()
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
			forestArt()
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

func healPlayer(player *data.Player) {
	if player.HP >= player.MaxHP {
	} else {
		player.HP = min(player.HP+15, player.MaxHP)
	}
	healDialog(player)
}

func playerAttack(player *data.Player, monster *data.Monster, monsterDefending *bool) {
	var weapon *data.Weapon
	for _, w := range data.Weapons {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range data.Musket {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range data.Lanter {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range data.Crossbow {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range data.Longsword {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	isRangedWeapon := player.WeaponID == 5 || player.WeaponID == 6 || player.WeaponID == 10 || player.WeaponID == 8

	if !isRangedWeapon && player.Position != monster.Position-1 {
		if data.Lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You're too far from %s to attack with your %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if data.Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти занадто далеко від %s щоб атакувати своїм %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты занадта далёка ад %s каб атакаваць сваім %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		}
		return
	}

	if weapon == nil {
		weapon = &data.Weapon{WeaponType: "Fists", Damage: 2, Stamina: 0}
	}

	playerDamage := player.Damage + rng.Rng()
	if *monsterDefending {
		blockUDialog()
		*monsterDefending = false
	} else if player.Stamina >= weapon.Stamina {
		player.Stamina -= weapon.Stamina
		if monster.ID == 17 {
			data.Tmp -= 1
			if data.Lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d steps to die", data.Tmp)).Foreground(termenv.ANSIBlue))
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d крокiв до смерті", data.Tmp)).Foreground(termenv.ANSIBlue))
			} else if data.Lang == "be" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d крокаў да мертвості", data.Tmp)).Foreground(termenv.ANSIBlue))
			}
		}
		if rng.Rng() == 1 {
			if data.Lang == "en" {
				fmt.Println(termenv.String("󰓥  Miss").Foreground(termenv.ANSIBlue))	
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String("󰓥  Мимо").Foreground(termenv.ANSIBlue))
			} else if data.Lang == "be" {
				fmt.Println(termenv.String("󰓥  Cпадарыня").Foreground(termenv.ANSIBlue))
			}
		} else {
		monster.HP -= playerDamage
		if player.WeaponID == 4 {
			monster.HP -= playerDamage
		}
		if data.Lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атакаваў %s на %d здароўя! Цяпер у яго %d ХП. Засталось %d вынослівасьці.", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти атакував %s з силою %d! Тепер в нього %d здоров'я. У тебе залишилось %d витривалостi", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		}}
	} else {
		noStaminaDialog()
	}
}

func playerSkip(player *data.Player) {
	if player.Stamina < 100 {
		player.Stamina = min(player.Stamina+20, player.MaxStamina)
		staminaDialog(player)
	} else {
		staminaDialog(player)
	}
}

func monsterTurnAction(monster *data.Monster, player *data.Player, monsterDefending *bool, playerDefending *bool, monsterAction string) {
    if monsterAction == "Defend" {
        if monster.Position < 5 {
            monster.Position += 1
            //if lang == "en" {
            //    fmt.Println(termenv.String(fmt.Sprintf("󰒙  The %s moves back!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            //} else if lang == "ua" {
            //    fmt.Println(termenv.String(fmt.Sprintf("󰒙  %s відступає!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            //} else if lang == "be" {
            //    fmt.Println(termenv.String(fmt.Sprintf("󰒙  %s адступае!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            //}
        }
        blockEnemyDialog()
        *monsterDefending = true
    } else if monsterAction == "Heal" {
        monster.HP = min(monster.HP+15, monster.MaxHP)
        healMonsterDialog(monster)
        *monsterDefending = false
    } else {
        if monster.Position > player.Position+1 {
            monster.Position -= 2
        }

        if monster.Position == player.Position+1 {
            monsterDamage := monster.Damage + rng.Rng() + monster.LVL
            if *playerDefending {
                blockEnemyAttack()
                *playerDefending = false
            } else {
                player.HP -= monsterDamage
                if data.Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s attacks you for %d damage! You now have %d HP.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else if data.Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s атакаваў цябе на %d здароўя! Цяпер у цябе %d ХП.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else {
		     fmt.Println(termenv.String(fmt.Sprintf("󰓥  Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.", 
                        monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
                }
            }
        } else {
            if data.Lang == "en" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s is too far away to attack!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if data.Lang == "ua" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадто далеко для атаки!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if data.Lang == "be" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадта далёка для атакі!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            }
        }
    }
}

func enemyTurn(monster *data.Monster, player *data.Player) string {
	if monster.ID == 10 {
		rngChoice := rng.Rng() % 3
		switch rngChoice {
		case 0: return "Attack"
		case 1: return "Defend"
		case 2: return "Heal"
		default: return "Attack"
	}
	} else {
		if player.Stamina < 10 {
			switch rng.Rng2() {
				case 1: return "Attack"
				case 2: return "Heal"
				default: return "Defend"}
		} else if monster.HP > 15 || player.HP < 15 {
			switch rng.Rng2() {
				case 1: return "Attack"
				case 2: return "Defend"
				default: return "Heal"}
		} else if monster.Position == player.Position-1 {
			switch rng.Rng2() {
				case 1: return "Defend"
				case 2: return "Attack"
				default: return "Heal"}
		} else if monster.HP < 15 {
			switch rng.Rng2() {
				case 1: return "Heal"
				case 2: return "Defend"
				default: return "Attack"}
		} else {
			return "Attack"
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

