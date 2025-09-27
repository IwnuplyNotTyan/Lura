package fight

import (
	"Lura/module/data"
	"Lura/module/dialog"
	"Lura/module/rng"
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

	if !isRangedWeapon && player.Position != monster.Position-1 && player.Position != monster.Position {
		dialog.TooFarDialog(player)
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
			dialog.GatoDialog()
		}
		if rng.Rng() == 1 {
			dialog.MissDialog()
		} else {
		monster.HP -= playerDamage
		if player.WeaponID == 4 {
			monster.HP -= playerDamage
		}
		dialog.PlayerAttackDialog(monster, player, playerDamage)
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
		dialog.AttackDialog(monster, player, monsterDamage)
            }
        } else {
		dialog.FarDialog(monster)
        }
    }
}

