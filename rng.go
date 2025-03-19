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
			"Upgrade Weapon",
			"Random Weapon",
			"Tears",
			"Broked heart",
			"Lotus",
			"Pearl necklace",
			"Turtle scute",
		}
	} else {
		buffs = []string{
			"Покращити зброю",
			"Випадкова зброя",
			"Розбите серце",
			"Черепаха щиткова",
			"Лотос",
			"Перлове намисто",
			"Сльози",
		}
	}
	return buffs[rand.Intn(len(buffs))]
}
