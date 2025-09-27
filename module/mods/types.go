package mods

import (
	"Lura/module/data"

	lua "github.com/yuin/gopher-lua"
)

func RegisterTypes(L *lua.LState) {
	mt := L.NewTypeMetatable("Monster")
	L.SetGlobal("Monster", mt)
	L.SetField(mt, "new", L.NewFunction(newMonster))
	L.SetField(mt, "setHP", L.NewFunction(setMonsterHP))
	L.SetField(mt, "getHP", L.NewFunction(getMonsterHP))
	L.SetField(mt, "remove", L.NewFunction(removeMonster))             
	L.SetField(mt, "removeByName", L.NewFunction(removeMonsterByName))

	// Register Weapon type
	wt := L.NewTypeMetatable("Weapon")
	L.SetGlobal("Weapon", wt)
	L.SetField(wt, "new", L.NewFunction(newWeapon))
	L.SetField(wt, "setDamage", L.NewFunction(setWeaponDamage))
	L.SetField(wt, "remove", L.NewFunction(removeWeapon))            
	L.SetField(wt, "removeByName", L.NewFunction(removeWeaponByName))

	// Expose global monsters and weapons tables to Lua
	monstersTable := L.NewTable()
	for _, monster := range data.Vmonsters {
		monsterTable := L.NewTable()
		L.SetField(monsterTable, "MonsterType", lua.LString(monster.MonsterType))
		L.SetField(monsterTable, "HP", lua.LNumber(monster.HP))
		L.SetField(monsterTable, "Damage", lua.LNumber(monster.Damage))
		L.SetField(monsterTable, "LVL", lua.LNumber(monster.LVL))
		L.SetField(monsterTable, "maxHP", lua.LNumber(monster.MaxHP))
		monstersTable.Append(monsterTable)
	}
	L.SetGlobal("monsters", monstersTable)

	weaponsTable := L.NewTable()
	for _, weapon := range data.Weapons {
		weaponTable := L.NewTable()
		L.SetField(weaponTable, "WeaponType", lua.LString(weapon.WeaponType))
		L.SetField(weaponTable, "Damage", lua.LNumber(weapon.Damage))
		L.SetField(weaponTable, "Stamina", lua.LNumber(weapon.Stamina))
		weaponsTable.Append(weaponTable)
	}
	L.SetGlobal("weapons", weaponsTable)
}
