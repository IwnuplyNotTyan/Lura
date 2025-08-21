package mods

import (
	"Lura/data"

	lua "github.com/yuin/gopher-lua"
)

func newMonster(L *lua.LState) int {
	monster := &data.Monster{
		MonsterType: L.CheckString(1),
		HP:          L.CheckInt(2),
		Damage:      L.CheckInt(3),
	}
	data.Vmonsters = append(data.Vmonsters, *monster)
	L.Push(lua.LNumber(len(data.Vmonsters) - 1))
	return 1
}

func setMonsterHP(L *lua.LState) int {
	idx := L.CheckInt(1)
	hp := L.CheckInt(2)
	data.Vmonsters[idx].HP = hp
	return 0
}

func getMonsterHP(L *lua.LState) int {
	idx := L.CheckInt(1)
	L.Push(lua.LNumber(data.Vmonsters[idx].HP))
	return 1
}

func newWeapon(L *lua.LState) int {
	weapon := &data.Weapon{
		WeaponType: L.CheckString(1),
		Damage:     L.CheckInt(2),
		Stamina:    L.CheckInt(3),
	}
	data.Weapons = append(data.Weapons, *weapon)
	//log.Printf("New weapon added: %+v", weapon)
	L.Push(lua.LNumber(len(data.Weapons) - 1))
	return 1
}

func setWeaponDamage(L *lua.LState) int {
	idx := L.CheckInt(1)
	damage := L.CheckInt(2)
	data.Weapons[idx].Damage = damage
	return 0
}

func removeMonster(L *lua.LState) int {
	idx := L.CheckInt(1)
	if idx < 0 || idx >= len(data.Vmonsters) {
		L.Push(lua.LString("Invalid monster index"))
		return 1
	}

	data.Vmonsters = append(data.Vmonsters[:idx], data.Vmonsters[idx+1:]...)
	L.Push(lua.LString("Monster removed successfully"))
	return 1
}

func removeWeapon(L *lua.LState) int {
	idx := L.CheckInt(1) // Get the index from Lua
	if idx < 0 || idx >= len(data.Weapons) {
		L.Push(lua.LString("Invalid weapon index"))
		return 1
	}

	data.Weapons = append(data.Weapons[:idx], data.Weapons[idx+1:]...)
	L.Push(lua.LString("Weapon removed successfully"))
	return 1
}

func removeMonsterByName(L *lua.LState) int {
	name := L.CheckString(1) // Get the name from Lua
	for i, monster := range data.Vmonsters {
		if monster.MonsterType == name {
			// Remove the monster from the slice
			data.Vmonsters = append(data.Vmonsters[:i], data.Vmonsters[i+1:]...)
			L.Push(lua.LString("Monster removed successfully"))
			return 1
		}
	}

	L.Push(lua.LString("Monster not found"))
	return 1
}

func removeWeaponByName(L *lua.LState) int {
	name := L.CheckString(1) // Get the name from Lua
	for i, weapon := range data.Weapons {
		if weapon.WeaponType == name {
			// Remove the weapon from the slice
			data.Weapons = append(data.Weapons[:i], data.Weapons[i+1:]...)
			L.Push(lua.LString("Weapon removed successfully"))
			return 1
		}
	}

	L.Push(lua.LString("Weapon not found"))
	return 1
}

