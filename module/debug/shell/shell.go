package debug

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Lura/module/data"
	"Lura/module/inv"
	"Lura/module/dialog"

	"github.com/muesli/termenv"
	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
)

func DebugShell(L *lua.LState, player *data.Player, monsters *data.Monster) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(termenv.String("  ").Foreground(termenv.ANSIBlue).Bold())
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			printHelp()
		case "clear":
			dialog.ClearScreen()
		case "checkAll":
			checkAll()
		case "checkAllDialog":
			checkkAllDialog(player, monsters)
		case "lang":
			if len(args) < 2 {
				log.Info("Usage: lang <language>")
				continue
			}
			data.Lang = os.Args[1]
		case "setScore":
  			if len(args) < 2 {
        			log.Info("Usage: setScore <value>")
        			continue
    			}
    			value, err := strconv.Atoi(args[1])
    			if err != nil {
    			    log.Info("Invalid value. Must be an integer.")
    			    continue
    			}
    			player.Score = value
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
		case "checkMods":
		        checkMods()

		case "seedData":
			data.SeedData()

		case "exit":
			fmt.Println("Exiting debug shell.")
			return
		case "listItem":
			inv.ShowInventory(player)
		case "addItem":
		    if len(args) < 5 {
		        fmt.Println("Usage: addItem <name> <effect> <value> <price>")
		        fmt.Println("Effects: heal, damage_boost, stamina_restore")
		        continue
		    }
	    
		    value, err := strconv.Atoi(args[3])
		    if err != nil {
		        fmt.Println("Invalid value. Must be an integer.")
		        continue
		    }
    
		    price, err := strconv.Atoi(args[4])
		    if err != nil {
		        fmt.Println("Invalid price. Must be an integer.")
		        continue
		    }
    
		    // Validate effect type
		    validEffects := map[string]bool{
		        "heal": true,
		        "damage_boost": true,
		        "stamina_restore": true,
			"material": true,
		    }
    
		    if !validEffects[args[2]] {
		        fmt.Println("Invalid effect. Valid effects are: material, heal, damage_boost, stamina_restore")
		        continue
		    }
    
		    inv.AddItem(&player.Inventory, args[1], args[2], value, price)
		    fmt.Printf("Added item: %s (Effect: %s, Value: %d, Price: %d)\n", 
		        args[1], args[2], value, price)
		case "checkPlayer":
			log.Infof("%+v", player)
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
				}
	}
}
