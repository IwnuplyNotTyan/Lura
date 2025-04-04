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
			if len(args) < 2 {
				fmt.Println("Usage: setHP <value>")
				continue
			}
			setHP(args[1], player)
		case "setLoc":
			if len(args) < 1 {
				fmt.Println("Usage: setLoc <value>")
				continue
			}
			setLoc(args[1], player)
		case "setDamage":
			if len(args) < 2 {
				fmt.Println("Usage: setDamage <value>")
				continue
			}
			setDamage(args[1], player)
		case "setHeart":
			if len(args) < 1 {
				fmt.Println("Usage: setHeart <value>")
				continue
			}
			setHeart(args[1], player)
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

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help                - Show this help message")
	fmt.Println("  checkAll            - List all monsters and weapons")
	fmt.Println("  setHP <value>       - Set HP of the player")
	fmt.Println("  setDamage <value>   - Set damage of the player")
	fmt.Println("  setLoc <value>      - Set location of the player")
	fmt.Println("  setHeart <value>    - Set heart of the player")
	fmt.Println("  addMonster <name> <hp> <damage> - Add a new monster")
	fmt.Println("  addWeapon <name> <damage> <stamina> - Add a new weapon")
	fmt.Println("  runLua <lua code>   - Execute Lua code")
	fmt.Println("  exit                - Exit the debug shell")
}

func setLoc(valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.loc = value
	fmt.Printf("Player location set to %d\n", value)
}

func setHeart(valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.heart = value
	fmt.Printf("Player heartbeat set to %d\n", value)
}

func setHP(valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.HP = value
	fmt.Printf("Player HP set to %d\n", value)
}

func setDamage(valueStr string, player *Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.Damage = value
	fmt.Printf("Player damage set to %d\n", value)
}

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

	vmonsters = append(vmonsters, Monster{MonsterType: name, HP: hp, Damage: damage})
	fmt.Printf("Added monster: %s (HP: %d, Damage: %d)\n", name, hp, damage)
}

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

func runLua(L *lua.LState, code string) {
	if err := L.DoString(code); err != nil {
		fmt.Printf("Lua error: %v\n", err)
	} else {
		fmt.Println("Lua code executed successfully.")
	}
}

func checkAllVWeapons() {
	fmt.Println(termenv.String("\n All V Weapons:").Foreground(termenv.ANSIBlue).Bold())

	for _, weapon := range weapons {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina Cost: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

	fmt.Println(termenv.String("\nDrops from monster:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range musket {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

	fmt.Println(termenv.String("\nFrom buffs:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range longsword {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}
	for _, weapon := range crossbow {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}
}

func checkAllSCWeapon() {
	fmt.Println(termenv.String("\n All SC Weapons:").Foreground(termenv.ANSIBlue).Bold())
	fmt.Println(termenv.String("\nDrops from monster:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range lanter {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

}
func checkAllVMonsters() {
	fmt.Println(termenv.String("\n All V Monsters:").Foreground(termenv.ANSIBlue).Bold())

	for _, monster := range vmonsters {
		fmt.Printf("Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
	}
}

func checkAllSCMonsters() {
	fmt.Println(termenv.String("\n All CC Monsters:").Foreground(termenv.ANSIBlue).Bold())

	for _, monster := range scmonsters {
		fmt.Printf("Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
	}
}

func checkAll() {
	checkAllVWeapons()
	checkAllSCWeapon()
	checkAllVMonsters()
	checkAllSCMonsters()
}
