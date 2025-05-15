package main

import (
	"math/rand"
)

func rng() int {
	return rand.Intn(6) + 1
}

func rng2() int {
	return rand.Intn(2) + 1
}

func rngHp() int {
    return rand.Intn(21) + 80
}

func getRandomWeapon() (string, int, int) {
	if len(weapons) == 0 {
		return "Fists", 2, 0
	}
	weapon := weapons[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage, weapon.ID
}

func getMusket(player *Player) {
	if len(musket) == 0 {
		return
	}
	weapon := musket[rand.Intn(len(musket))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * rng()
	player.WeaponID = weapon.ID
}

func getLanter(player *Player) {
	if len(lanter) == 0 {
		return
	}
	weapon := lanter[rand.Intn(len(lanter))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * rng()
	player.WeaponID = weapon.ID
}

func getCrossbow(player *Player) {
	weapon := crossbow[rand.Intn(len(crossbow))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * rng()
	player.WeaponID = weapon.ID
}

func getLongsword(player *Player) {
	weapon := longsword[rand.Intn(len(longsword))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * rng()
	player.WeaponID = weapon.ID
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

func getRandomBoss() *Monster {
	if len(boss) == 0 {
		return nil
	}
	monster := boss[rand.Intn(len(boss))]
	monster.LVL = rand.Intn(5) + 1
	monster.maxHP = monster.HP + (monster.LVL * 10)
	return &monster
}
