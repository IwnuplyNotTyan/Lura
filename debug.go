package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/muesli/termenv"
	lua "github.com/yuin/gopher-lua"
)

// DebugShell starts an interactive debug shell.
func DebugShell(L *lua.LState, player *Player) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Debug Shell started. Type 'help' for a list of commands.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			printHelp()
		case "checkAll":
			checkAll()
		case "setHP":
			if len(args) < 3 {
				fmt.Println("Usage: setHP <monster|player> <value>")
				continue
			}
			setHP(args[1], args[2], player)
		case "setDamage":
			if len(args) < 3 {
				fmt.Println("Usage: setDamage <monster|player> <value>")
				continue
			}
			setDamage(args[1], args[2], player)
		case "addMonster":
			if len(args) < 4 {
				fmt.Println("Usage: addMonster <name> <hp> <damage>")
				continue
			}
			addMonster(args[1], args[2], args[3])
		case "addWeapon":
			if len(args) < 4 {
				fmt.Println("Usage: addWeapon <name> <damage> <stamina>")
				continue
			}
			addWeapon(args[1], args[2], args[3])
		case "runLua":
			if len(args) < 2 {
				fmt.Println("Usage: runLua <lua code>")
				continue
			}
			runLua(L, strings.Join(args[1:], " "))
		case "exit":
			fmt.Println("Exiting debug shell.")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}

// printHelp displays a list of available debug commands.
func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help                - Show this help message")
	fmt.Println("  checkAll            - List all monsters and weapons")
	fmt.Println("  setHP <target> <value> - Set HP of a monster or player")
	fmt.Println("  setDamage <target> <value> - Set damage of a monster or player")
	fmt.Println("  addMonster <name> <hp> <damage> - Add a new monster")
	fmt.Println("  addWeapon <name> <damage> <stamina> - Add a new weapon")
	fmt.Println("  runLua <lua code>   - Execute Lua code")
	fmt.Println("  exit                - Exit the debug shell")
}

// setHP sets the HP of a monster or player.
func setHP(target string, valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	switch target {
	case "player":
		player.HP = value
		fmt.Printf("Player HP set to %d\n", value)
	case "monster":
		if len(monsters) == 0 {
			fmt.Println("No monsters available.")
			return
		}
		monsters[0].HP = value // Modify the first monster for simplicity
		fmt.Printf("Monster HP set to %d\n", value)
	default:
		fmt.Println("Invalid target. Use 'player' or 'monster'.")
	}
}

// setDamage sets the damage of a monster or player.
func setDamage(target string, valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	switch target {
	case "player":
		player.Damage = value
		fmt.Printf("Player damage set to %d\n", value)
	case "monster":
		if len(monsters) == 0 {
			fmt.Println("No monsters available.")
			return
		}
		monsters[0].Damage = value // Modify the first monster for simplicity
		fmt.Printf("Monster damage set to %d\n", value)
	default:
		fmt.Println("Invalid target. Use 'player' or 'monster'.")
	}
}

// addMonster adds a new monster to the game.
func addMonster(name, hpStr, damageStr string) {
	hp, err := strconv.Atoi(hpStr)
	if err != nil {
		fmt.Println("Invalid HP. Must be an integer.")
		return
	}

	damage, err := strconv.Atoi(damageStr)
	if err != nil {
		fmt.Println("Invalid damage. Must be an integer.")
		return
	}

	monsters = append(monsters, Monster{MonsterType: name, HP: hp, Damage: damage})
	fmt.Printf("Added monster: %s (HP: %d, Damage: %d)\n", name, hp, damage)
}

// addWeapon adds a new weapon to the game.
func addWeapon(name, damageStr, staminaStr string) {
	damage, err := strconv.Atoi(damageStr)
	if err != nil {
		fmt.Println("Invalid damage. Must be an integer.")
		return
	}

	stamina, err := strconv.Atoi(staminaStr)
	if err != nil {
		fmt.Println("Invalid stamina. Must be an integer.")
		return
	}

	weapons = append(weapons, Weapon{WeaponType: name, Damage: damage, Stamina: stamina})
	fmt.Printf("Added weapon: %s (Damage: %d, Stamina: %d)\n", name, damage, stamina)
}

// runLua executes Lua code from the debug shell.
func runLua(L *lua.LState, code string) {
	if err := L.DoString(code); err != nil {
		fmt.Printf("Lua error: %v\n", err)
	} else {
		fmt.Println("Lua code executed successfully.")
	}
}

// checkAllWeapons lists all weapons.
func checkAllWeapons() {
	if lang == "en" {
		fmt.Println(termenv.String("\n All Weapons:").Foreground(termenv.ANSIBlue).Bold())
	} else {
		fmt.Println(termenv.String("\n Вся зброя:").Foreground(termenv.ANSIBlue).Bold())
	}

	for _, weapon := range weapons {
		if lang == "en" {
			fmt.Printf(" Weapon: %s, Damage: %d, Stamina Cost: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
		} else {
			fmt.Printf(" Зброя: %s, Пошкодження: %d, Витрати витривалості: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
		}
	}
}

// checkAllMonsters lists all monsters.
func checkAllMonsters() {
	if lang == "en" {
		fmt.Println(termenv.String("\n All Monsters:").Foreground(termenv.ANSIBlue).Bold())
	} else {
		fmt.Println(termenv.String("\n Всі монстри:").Foreground(termenv.ANSIBlue).Bold())
	}

	for _, monster := range monsters {
		if lang == "en" {
			fmt.Printf(" Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
		} else {
			fmt.Printf(" Монстр: %s, Здоров'я: %d, Пошкодження: %d, Рівень: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
		}
	}
}

// checkAll lists all weapons and monsters.
func checkAll() {
	checkAllWeapons()
	checkAllMonsters()
}
