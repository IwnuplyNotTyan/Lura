package main

import (
	"fmt"
	"strings"
	"time"

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

func displayPositions(player *Player, monster *Monster) {
    positions := make([]string, 6)
    for i := range positions {
        positions[i] = " "
    }
    
    positions[player.Position] = " "
    positions[monster.Position] = " "
    
    fmt.Println(strings.Join(positions, ""))
}

func takeWeapon(player *Player, monster *Monster) {
	var confirm bool
	var a, b, c string

	if lang == "en" {
		a = "Do you want to take the weapon?"
		b = "Yes"
		c = "No"
	} else if lang == "ua" {
		a = "Ви хочете взяти зброю?"
		b = "Так"
		c = "Ні"
	} else if lang == "be" {
		a = "Вы хочаце ўзяць зброю?"
		b = "Так"
		c = "Не"
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
	log.Info(confirm)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirm {
		if monster.MonsterType == "Lanter keeper" || monster.MonsterType == "Ахоўца ліхтара" || monster.MonsterType == "Охоронець ліхтаря" {
			getLanter(player)
		} else if monster.ID == 1 {
			getMusket(player)
		}
	}
}

func afterLoc(player *Player) {
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
	staminaSleep := player.maxStamina + 20
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

func crafting(player *Player) {
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

func selectAttack() string {
	var selectedAttack string

	if lang == "ua" {
		Attack = "Атакувати"
		Heal = "Лікуватися"
		Defend = "Захищатися"
		Skip = "Пропустити"
	} else if lang == "be" {
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

	f := huh.NewForm(
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

	if err := f.Run(); err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	clearScreen()
	return selectedAttack
}

func fight(player *Player, monster *Monster, config *Config, weapon *Weapon) {
	//afterLoc(player)
	for player.HP > 0 {
		//log.Info(player.score)
		if player.loc == 0 {
			monster = getRandomCMonster()
		} else if player.loc == 1 {
			monster = getRandomVMonster()
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
			playerAction := selectAttack()

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
						if lang == "en" {
							fmt.Println(termenv.String(fmt.Sprintf("󰓥  You not so close to %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
						} else if lang == "ua" {
							fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти не так близько до %s", monster.MonsterType)).Foreground(termenv.ANSIBrightRed))
						} else if lang == "be" {
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

			if player.HP <= 0 {
				if player.heart == 1 {
					if player.score > config.Score {
						config.Score = player.score
						if lang == "en" {
							fmt.Println(termenv.String(fmt.Sprintf("  New High Score, %d", player.score)).Foreground(termenv.ANSIBrightRed).Bold())
						} else if lang == "ua" {
							fmt.Println(termenv.String(fmt.Sprintf("  Новий рекорд, %d", player.score)).Foreground(termenv.ANSIBrightRed).Bold())
						} else if lang == "be" {
							fmt.Println(termenv.String(fmt.Sprintf("  Новы рэкорд, %d", player.score)).Foreground(termenv.ANSIBrightRed).Bold())
						}
						if err := saveConfig(getConfigPath(), *config); err != nil {
							log.Printf("Error saving high score: %v", err)
						}
					} else {
						fmt.Println(termenv.String(fmt.Sprintf("  %d", player.score)).Foreground(termenv.ANSIRed).Bold())
					}
					return
				} else if player.heart == 0 {
					player.maxHP = player.maxHP / 2
					player.HP = player.maxHP
					player.Damage = player.Damage * 2
					if lang == "en" {
						fmt.Println(termenv.String(fmt.Sprintf("  Your heart is broken! HP set to %d, Damage increased to %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					} else if lang == "ua" {
						fmt.Println(termenv.String(fmt.Sprintf("  Ваше серце розбито! HP встановлено на %d, Пошкодження збільшено до %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					} else if lang == "be" {
						fmt.Println(termenv.String(fmt.Sprintf("  Ваша сэрца разбіта! HP устаноўлена на %d, Пашкоджанні павялічаны да %d.", player.HP, player.Damage)).Foreground(termenv.ANSIBrightRed).Bold())
					}
					player.heart = 1
					fight(player, monster, config, weapon)
				}
			}
			time.Sleep(time.Second)
		}

		clearScreen()
		player.score += monster.score
		player.Coins += monster.coins

		player.Position = 0
		monster.Position = 5

		defeatMonster(monster)
		if monster.ID == 2 {
			if rng2() == 1 {
				takeWeapon(player, monster)
			}
		} else if monster.ID == 1 {
			if rng2() == 1 {
				takeWeapon(player, monster)
			}
		}
		if player.WeaponType == "Lanter of the soul" {
			if player.time == 1 {
				player.Damage = monster.Damage * rng()
				player.monster = true
				player.HP = monster.maxHP
				player.maxHP = monster.maxHP
				player.name = monster.MonsterType
				player.WeaponType = ""

				if lang == "en" {
					fmt.Println(termenv.String(fmt.Sprintf("  Now you %s", player.name)).Foreground(termenv.ANSIRed).Bold())
				} else if lang == "ua" {
					fmt.Println(termenv.String(fmt.Sprintf("  Тепер ти %s", player.name)).Foreground(termenv.ANSIRed).Bold())
				} else if lang == "be" {
					fmt.Println(termenv.String(fmt.Sprintf("  Цяпер ты %s", player.name)).Foreground(termenv.ANSIRed).Bold())
				}
			} else {
				player.time = 1
			}
		}

		if player.buffs == 4 {
			if player.monster == false {
				buffsAction(player)
			}
			player.buffs = 0
			newLine()
			if player.loc == 0 {
				forestArt()
				player.loc = 1
			} else if player.loc == 1 {
				caveArt()
				player.loc = 0
			}
		} else {
			player.buffs += 1
			if lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Step to another location", player.buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			} else if lang == "ua" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Крокiв до iншого мiсця", player.buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			} else if lang == "be" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d Крокаў да баффу", player.buffs)).Foreground(termenv.ANSIBrightMagenta).Bold())
			}
		}
	}

	if player.heart == 2 {
		player.HP += 10
	}
}

func healPlayer(player *Player) {
	player.HP = min(player.HP+15, player.maxHP)
	healDialog(player)
}

func playerAttack(player *Player, monster *Monster, monsterDefending *bool) {
	var weapon *Weapon
	for _, w := range weapons {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range musket {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range lanter {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range crossbow {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	for _, w := range longsword {
		if w.WeaponType == player.WeaponType {
			weapon = &w
			break
		}
	}

	isRangedWeapon := player.WeaponID == 5 || player.WeaponID == 6 || player.WeaponID == 10 || player.WeaponID == 8

	if !isRangedWeapon && player.Position != monster.Position-1 {
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You're too far from %s to attack with your %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти занадто далеко від %s щоб атакувати своїм %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты занадта далёка ад %s каб атакаваць сваім %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		}
		return
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
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else if lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атакаваў %s на %d здароўя! Цяпер у яго %d ХП. Засталось %d вынослівасьці.", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти атакував %s з силою %d! Тепер в нього %d здоров'я. У тебе залишилось %d витривалостi", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		}
	} else {
		noStaminaDialog()
	}
}

func playerSkip(player *Player) {
	if player.Stamina < 100 {
		player.Stamina = min(player.Stamina+20, player.maxStamina)
		staminaDialog(player)
	} else {
		staminaDialog(player)
	}
}

func monsterTurnAction(monster *Monster, player *Player, monsterDefending *bool, playerDefending *bool, monsterAction string) {
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
        monster.HP = min(monster.HP+15, monster.maxHP)
        healMonsterDialog(monster)
        *monsterDefending = false
    } else {
        if monster.Position > player.Position+1 {
            monster.Position -= 2
        }

        if monster.Position == player.Position+1 {
            monsterDamage := monster.Damage + rng() + monster.LVL
            if *playerDefending {
                blockEnemyAttack()
                *playerDefending = false
            } else {
                player.HP -= monsterDamage
                if lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s attacks you for %d damage! You now have %d HP.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else if lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s атакаваў цябе на %d здароўя! Цяпер у цябе %d ХП.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else {
		     fmt.Println(termenv.String(fmt.Sprintf("󰓥  Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.", 
                        monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
                }
            }
        } else {
            if lang == "en" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s is too far away to attack!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if lang == "ua" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадто далеко для атаки!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if lang == "be" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадта далёка для атакі!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            }
        }
    }
}

/*func enemyTurn(monster *Monster) string {
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
}*/

func enemyTurn(monster *Monster, player *Player) string {
	if monster.HP > 15 || player.HP < 15 {
		return "Attack"
	} else if monster.Position == player.Position+1 {
		return "Defend"
	} else if monster.HP < 15 {
		return "Heal"
	} else {
		return "Attack"
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

