package debug

import (
	"fmt"
	"strconv"

	"Lura/module/data"
	"Lura/module/dialog"
	"Lura/module/mods"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/log"
	"github.com/muesli/termenv"
	lua "github.com/yuin/gopher-lua"
)

func printHelp() {
	help := `| Command | Description |
| ------- | ----------- |
| help | Show this help message |
| checkAll | List all monsters and weapons |
| setHP \<value\> | Set HP of the player |
| setDamage \<value\> | Set damage of the player |
| setLoc \<value\> | Set location of the player |
| setHeart \<value\> | Set heart of the player |
| setScore \<value\> | Set score of the player |
| addMonster \<name\> \<hp\> \<damage\> | Add a new monster |
| addWeapon \<name\> \<damage\> \<stamina\> | Add a new weapon |
| AddItem \<name\> \<effect\> \<value\> \<price\> | Add a new item to the inventory |
| listItem | List all items in the inventory |
| checkMods | Check loaded mods |
| checkPlayer | Check player stats |
| seedData | Seed data for testing |
| runLua \<lua code\> | Execute Lua code |
| clear | Clear terminal logs |
| exit | Exit the debug shell |`
	out, err := glamour.Render(help, "dark")
	if err != nil {
		log.Info("Error rendering help:", err)
	}
	fmt.Println(out)
}

func checkMods() {
    mods := mods.GetLoadedMods()
    if len(mods) == 0 {
        log.Info("No mods loaded")
        return
    }
    
    fmt.Println("Loaded mods:")
    for i, mod := range mods {
        log.Info("%d. %s\n", i+1, mod)
    }
}

func setLoc(valueStr string, player *data.Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Info("Invalid value. Must be an integer.")
		return
	}
	player.Loc = value
	fmt.Printf("Player location set to %d\n", value)
}

func setHeart(valueStr string, player *data.Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.Heart = value
	fmt.Printf("Player heartbeat set to %d\n", value)
}

func setHP(valueStr string, player *data.Player) {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Must be an integer.")
		return
	}

	player.HP = value
	fmt.Printf("Player HP set to %d\n", value)
}

func setDamage(valueStr string, player *data.Player) {
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

	data.Vmonsters = append(data.Vmonsters, data.Monster{MonsterType: name, HP: hp, Damage: damage})
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

	data.Weapons = append(data.Weapons, data.Weapon{WeaponType: name, Damage: damage, Stamina: stamina})
	fmt.Printf("Added weapon: %s (Damage: %d, Stamina: %d)\n", name, damage, stamina)
}

func runLua(L *lua.LState, code string) {
	if err := L.DoString(code); err != nil {
		fmt.Printf("Lua error: %v\n", err)
	} else {
		fmt.Println("Lua code executed successfully.")
	}
}

func checkkAllDialog(player *data.Player, monsters *data.Monster) {
	fmt.Println(termenv.String("\n  All Dialogs:").Foreground(termenv.ANSIBlue).Bold())
	fmt.Println(termenv.String("\n%s:", data.Lang).Foreground(termenv.ANSICyan).Italic())
	dialog.BlockDialog()
	dialog.BlockEnemyAttack()
	dialog.BlockEnemyDialog()
	dialog.BlockUDialog()
	dialog.DefeatMonster(monsters)
	dialog.HealDialog(player)
	dialog.HealMonsterDialog(monsters)
	dialog.NoBuffDialog()
	dialog.NoStaminaDialog()
	dialog.StaminaDialog(player)
	dialog.NoBuffDialog()
}

func checkAllVWeapons() {
	fmt.Println(termenv.String("\n  All V Weapons:").Foreground(termenv.ANSIBlue).Bold())

	for _, weapon := range data.Weapons {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina Cost: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

	fmt.Println(termenv.String("\nDrops from monster:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range data.Musket {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

	fmt.Println(termenv.String("\nFrom buffs:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range data.Longsword {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}
	for _, weapon := range data.Crossbow {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}
}

func checkAllSCWeapon() {
	fmt.Println(termenv.String("\n  All SC Weapons:").Foreground(termenv.ANSIBlue).Bold())
	fmt.Println(termenv.String("\nDrops from monster:").Foreground(termenv.ANSICyan).Italic())
	for _, weapon := range data.Lanter {
		fmt.Printf("Weapon: %s, Damage: %d, Stamina: %d\n", weapon.WeaponType, weapon.Damage, weapon.Stamina)
	}

}
func checkAllVMonsters() {
	fmt.Println(termenv.String("\n  All V Monsters:").Foreground(termenv.ANSIBlue).Bold())

	for _, monster := range data.Vmonsters {
		fmt.Printf("Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
	}
}

func checkAllSCMonsters() {
	fmt.Println(termenv.String("\n  All CC Monsters:").Foreground(termenv.ANSIBlue).Bold())

	for _, monster := range data.Scmonsters {
		fmt.Printf("Monster: %s, HP: %d, Damage: %d, Level: %d\n", monster.MonsterType, monster.HP, monster.Damage, monster.LVL)
	}
}

func checkAll() {
	checkAllVWeapons()
	checkAllSCWeapon()
	checkAllVMonsters()
	checkAllSCMonsters()
}
