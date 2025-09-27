package rng

import (
	"math/rand"

	"Lura/module/data"
)

func GetRandomWeapon() (string, int, int) {
	if len(data.Weapons) == 0 {
		return "Error", 0, 0
	}
	weapon := data.Weapons[rand.Intn(len(data.Weapons))]
	return weapon.WeaponType, weapon.Damage, weapon.ID
}

func GetMusket(player *data.Player) {
	if len(data.Musket) == 0 {
		return
	}
	weapon := data.Musket[rand.Intn(len(data.Musket))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * Rng()
	player.WeaponID = weapon.ID
}

func GetLanter(player *data.Player) {
	if len(data.Lanter) == 0 {
		return
	}
	weapon := data.Lanter[rand.Intn(len(data.Lanter))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * Rng()
	player.WeaponID = weapon.ID
}

func GetCrossbow(player *data.Player) {
	weapon := data.Crossbow[rand.Intn(len(data.Crossbow))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * Rng()
	player.WeaponID = weapon.ID
}

func GetLongsword(player *data.Player) {
	weapon := data.Longsword[rand.Intn(len(data.Longsword))]
	player.WeaponType = weapon.WeaponType
	player.Damage = weapon.Damage * Rng()
	player.WeaponID = weapon.ID
}

func GetRandomVMonster() *data.Monster {
	if len(data.Vmonsters) == 0 {
		return nil
	}
	monster := data.Vmonsters[rand.Intn(len(data.Vmonsters))]
	monster.LVL = rand.Intn(5) + 1
	monster.MaxHP = monster.HP + (monster.LVL * 10)
	return &monster
}

func GetRandomCMonster() *data.Monster {
	if len(data.Scmonsters) == 0 {
		return nil
	}
	monster := data.Scmonsters[rand.Intn(len(data.Scmonsters))]
	monster.LVL = rand.Intn(5) + 1
	monster.MaxHP = monster.HP + (monster.LVL * 10)
	return &monster
}

func GetRandomBoss() *data.Monster {
	if len(data.Boss) == 0 {
		return nil
	}
	monster := data.Boss[rand.Intn(len(data.Boss))]
	monster.LVL = rand.Intn(5) + 1
	monster.MaxHP = monster.HP + (monster.LVL * 10)
	return &monster
}
