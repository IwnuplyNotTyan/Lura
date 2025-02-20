package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

func loadLuaScript(L *lua.LState, scriptPath string) error {
	return L.DoFile(scriptPath)
}

func AutoLoadMods(L *lua.LState) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	modsDir := filepath.Join(homeDir, ".local", "share", "Lura", "mods")

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
		}

		return nil
	})

	return err
}

func registerTypes(L *lua.LState) {
	mt := L.NewTypeMetatable("Monster")
	L.SetGlobal("Monster", mt)
	L.SetField(mt, "new", L.NewFunction(newMonster))
	L.SetField(mt, "setHP", L.NewFunction(setMonsterHP))
	L.SetField(mt, "getHP", L.NewFunction(getMonsterHP))

	wt := L.NewTypeMetatable("Weapon")
	L.SetGlobal("Weapon", wt)
	L.SetField(wt, "new", L.NewFunction(newWeapon))
	L.SetField(wt, "setDamage", L.NewFunction(setWeaponDamage))
}

func newMonster(L *lua.LState) int {
	monster := &Monster{
		MonsterType: L.CheckString(1),
		HP:          L.CheckInt(2),
		Damage:      L.CheckInt(3),
	}
	monsters = append(monsters, *monster)
	L.Push(lua.LNumber(len(monsters) - 1))
	return 1
}

func setMonsterHP(L *lua.LState) int {
	idx := L.CheckInt(1)
	hp := L.CheckInt(2)
	monsters[idx].HP = hp
	return 0
}

func getMonsterHP(L *lua.LState) int {
	idx := L.CheckInt(1)
	L.Push(lua.LNumber(monsters[idx].HP))
	return 1
}

func newWeapon(L *lua.LState) int {
	weapon := &Weapon{
		WeaponType: L.CheckString(1),
		Damage:     L.CheckInt(2),
		Stamina:    L.CheckInt(3),
	}
	weapons = append(weapons, *weapon)
	//log.Printf("New weapon added: %+v", weapon)
	L.Push(lua.LNumber(len(weapons) - 1))
	return 1
}

func setWeaponDamage(L *lua.LState) int {
	idx := L.CheckInt(1)
	damage := L.CheckInt(2)
	weapons[idx].Damage = damage
	return 0
}
