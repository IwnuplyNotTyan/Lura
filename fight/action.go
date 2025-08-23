package fight

import (
	"fmt"

	"Lura/data"
	"Lura/module/dialog"
	"Lura/module/rng"

	"github.com/muesli/termenv"
)

func healPlayer(player *data.Player) {
	if player.HP >= player.MaxHP {
	} else {
		player.HP = min(player.HP+15, player.MaxHP)
	}
	dialog.HealDialog(player)
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
		if data.Lang == "ru" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты слишком далеко от %s чтобы атаковать своим %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if data.Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти занадто далеко від %s щоб атакувати своїм %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты занадта далёка ад %s каб атакаваць сваім %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You're too far from %s to attack with your %s!", monster.MonsterType, player.WeaponType)).Foreground(termenv.ANSIYellow))
		}
		return
	}

	if weapon == nil {
		weapon = &data.Weapon{WeaponType: "Fists", Damage: 2, Stamina: 0}
	}

	playerDamage := player.Damage + rng.Rng()
	if *monsterDefending {
		dialog.BlockUDialog()
		*monsterDefending = false
	} else if player.Stamina >= weapon.Stamina {
		player.Stamina -= weapon.Stamina
		if monster.ID == 17 {
			data.Tmp -= 1
			if data.Lang == "ru" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d шагов до смерти", data.Tmp)).Foreground(termenv.ANSIBlue))
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d крокiв до смерті", data.Tmp)).Foreground(termenv.ANSIBlue))
			} else if data.Lang == "be" {
				fmt.Println(termenv.String(fmt.Sprintf("  %d крокаў да мертвості", data.Tmp)).Foreground(termenv.ANSIBlue))
			} else {
				fmt.Println(termenv.String(fmt.Sprintf("  %d steps to die", data.Tmp)).Foreground(termenv.ANSIBlue))
			}
		}
		if rng.Rng() == 1 {
			if data.Lang == "ru" {
				fmt.Println(termenv.String("󰓥  Мимо").Foreground(termenv.ANSIBlue))	
			} else if data.Lang == "ua" {
				fmt.Println(termenv.String("󰓥  Мимо").Foreground(termenv.ANSIBlue))
			} else if data.Lang == "be" {
				fmt.Println(termenv.String("󰓥  Cпадарыня").Foreground(termenv.ANSIBlue))
			} else {
				fmt.Println(termenv.String("󰓥  Miss").Foreground(termenv.ANSIBlue))	
			}
		} else {
		monster.HP -= playerDamage
		if player.WeaponID == 4 {
			monster.HP -= playerDamage
		}
		if data.Lang == "ru" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ты атакаваў %s на %d здароўя! Цяпер у яго %d ХП. Засталось %d вынослівасьці.", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else if data.Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  Ти атакував %s з силою %d! Тепер в нього %d здоров'я. У тебе залишилось %d витривалостi", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥  You attack the %s for %d damage! It now has %d HP. %d stamina remaining", monster.MonsterType, playerDamage, monster.HP, player.Stamina)).Foreground(termenv.ANSIBlue))
		}
	}
	} else {
		dialog.NoStaminaDialog()
	}
}

func playerSkip(player *data.Player) {
	if player.Stamina < 100 {
		player.Stamina = min(player.Stamina+20, player.MaxStamina)
		dialog.StaminaDialog(player)
	} else {
		dialog.StaminaDialog(player)
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
        dialog.BlockEnemyDialog()
        *monsterDefending = true
    } else if monsterAction == "Heal" {
        monster.HP = min(monster.HP+15, monster.MaxHP)
        dialog.HealMonsterDialog(monster)
        *monsterDefending = false
    } else {
        if monster.Position > player.Position+1 {
            monster.Position -= 2
        }

        if monster.Position == player.Position+1 {
            monsterDamage := monster.Damage + rng.Rng() + monster.LVL
            if *playerDefending {
                dialog.BlockEnemyAttack()
                *playerDefending = false
            } else {
                player.HP -= monsterDamage
                if data.Lang == "ru" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s attacks you for %d damage! You now have %d HP.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else if data.Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s атакаваў цябе на %d здароўя! Цяпер у цябе %d ХП.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
                } else if data.Lang == "ua" {
		     fmt.Println(termenv.String(fmt.Sprintf("󰓥  Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.", 
                        monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
                } else {
                    fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s attacks you for %d damage! You now have %d HP.", 
                        monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
		}
            }
        } else {
            if data.Lang == "ru" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s is too far away to attack!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if data.Lang == "ua" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадто далеко для атаки!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else if data.Lang == "be" {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s занадта далёка для атакі!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
            } else {
                fmt.Println(termenv.String(fmt.Sprintf("󰓥  The %s is too far away to attack!", monster.MonsterType)).Foreground(termenv.ANSIYellow))
	    }
        }
    }
}

