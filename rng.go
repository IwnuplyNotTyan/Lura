package main

import (
	"log"
	"math/rand"
)

func rng() int {
	return rand.Intn(6) + 1
}

func rng2() int {
	return rand.Intn(2) + 1
}

func getRandomWeapon() (string, int) {
	if len(weapons) == 0 {
		log.Println("Warning: No weapons available. Using default weapon (Fists, 2).")
		return "Fists", 2
	}
	weapon := weapons[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getMusket() (string, int) {
	weapon := musket[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getLanter() (string, int) {
	weapon := lanter[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getCrossbow() (string, int) {
	weapon := crossbow[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getLongsword() (string, int) {
	weapon := longsword[rand.Intn(len(weapons))]
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

func getRandomCMonster() *Monster {
	if len(scmonsters) == 0 {
		return nil
	}
	monster := scmonsters[rand.Intn(len(scmonsters))]
	monster.LVL = rand.Intn(5) + 1
	monster.maxHP = monster.HP + (monster.LVL * 10)
	return &monster
}
