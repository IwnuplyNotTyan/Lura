package main

import (
	"fmt"
	"os"
	"path/filepath"

	"Lura/data"

	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
)

var loadedMods []string

func AutoLoadMods(L *lua.LState) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	modsDir := filepath.Join(homeDir, ".config", "lura", "mods")
	loadedMods = []string{}

	if err := os.MkdirAll(modsDir, 0755); err != nil {
		return fmt.Errorf("failed to create mods directory: %v", err)
	}

	err = filepath.Walk(modsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".lua" {
			return nil
		}

		if err := L.DoFile(path); err != nil {
			log.Printf("Failed to load mod %s: %v", path, err)
		} else {
			loadedMods = append(loadedMods, filepath.Base(path))
		}

		return nil
	})

	return err
}

func ModsLoaded() bool {
    return len(loadedMods) > 0
}

func GetLoadedMods() []string {
	return loadedMods
}

func registerTypes(L *lua.LState) {
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

