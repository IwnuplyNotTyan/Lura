package main

import (
	"log"
	"math/rand"
)

func rng() int {
	return rand.Intn(6) + 1
}

func getRandomWeapon() (string, int) {
	if len(weapons) == 0 {
		log.Println("Warning: No weapons available. Using default weapon (Fists, 2).")
		return "Fists", 2
	}
	weapon := weapons[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getRandomVMonster() *Monster {
	if len(vmonsters) == 0 {
		return nil
	}
	monster := vmonsters[rand.Intn(len(vmonsters))]
	monster.LVL = rand.Intn(5) + 1
	monster.maxHP = monster.HP + (monster.LVL * 10)
	return &monster
}

//func getRandomSMonster() *Monster {
//	if len(smonsters) == 0 {
//		return nil
//	}
//	monster := smonsters[rand.Intn(len(smonsters))]
//	monster.LVL = rand.Intn(5) + 1
//	monster.maxHP = monster.HP + (monster.LVL * 10)
//	return &monster
//}

//func getRandomCCMonster() *Monster {
//	if len(ccmonsters) == 0 {
//		return nil
//	}
//	monster := ccmonsters[rand.Intn(len(ccmonsters))]
//	monster.LVL = rand.Intn(5) + 1
//	monster.maxHP = monster.HP + (monster.LVL * 10)
//	return &monster
//}

func getRandomBuff() string {
	var buffs []string

	if lang == "en" {
		buffs = []string{
			"Increase HP (+2) & Reduce Damage (-1)",
			"Increase Damage (+5) & Reduce HP (-5)",
			"Add Armor (+50)",
			"Upgrade Weapon",
			"Increase Stamina (+10) & Reduce Damage (-2)",
			"Random Weapon",
		}
	} else {
		buffs = []string{
			"Додано здоров'я (+2) & Зменшено пошкодження (-1)",
			"Додано пошкодження (+5) & Зменшено здоров'я (-5)",
			"Добавити захисту (+50)",
			"Покращити зброю",
			"Додано витривалiсть (+10) & Зменшино пошкодження (-2)",
			"Випадкова зброя",
		}
	}
	return buffs[rand.Intn(len(buffs))]
}
