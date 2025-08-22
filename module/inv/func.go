package inv

import (
	"fmt"

	"Lura/data"

	"github.com/muesli/termenv"
)

func AddItem(inv *data.Inventory, name string, effect string, value int, price int) {
    // Check if item already exists
    for i, item := range inv.Items {
        if item.Name == name {
            inv.Items[i].Quantity++
            return
        }
    }
    
    // Add new item
    newItem := data.Item{
        ID:       inv.NextID,
        Name:     name,
        Quantity: 1,
        Effect:   effect,
        Value:    value,
        Price:    price,
    }
    inv.Items = append(inv.Items, newItem)
    inv.NextID++
}

func RemoveItem(inv *data.Inventory, id int, quantity int) bool {
    for i, item := range inv.Items {
        if item.ID == id {
            if item.Quantity > quantity {
                inv.Items[i].Quantity -= quantity
            } else {
                // Remove item completely if quantity is 0 or negative
                inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
            }
            return true
        }
    }
    return false
}

func UseItem(player *data.Player, id int) {
    for _, item := range player.Inventory.Items {
        if item.ID == id {
            switch item.Effect {
	    case "Material":
		    if data.Lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  You get %s, +%d material!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if data.Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  Ви отримали %s, +%d матеріалів!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if data.Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Вы атрымалі %s, +%d матэрыялаў!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    }
            case "heal":
                player.HP = min(player.HP+item.Value, player.MaxHP)
                if data.Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, healed %d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "damage_boost":
                player.Damage += item.Value
                if data.Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, +%d damage for this fight!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d пошкоджень у цьому бою!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d пашкоджанняў у гэтым баі!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "stamina_restore":
                player.Stamina = min(player.Stamina+item.Value, player.MaxStamina)
                if data.Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, restored %d stamina!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d витривалостi!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if data.Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d вынослівасці!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
            }
            
            RemoveItem(&player.Inventory, id, 1)
            return
        }
    }
    
    if data.Lang == "en" {
        fmt.Println(termenv.String("  Item not found!").Foreground(termenv.ANSIRed))
    } else if data.Lang == "ua" {
        fmt.Println(termenv.String("  Предмет не знайдено!").Foreground(termenv.ANSIRed))
    } else if data.Lang == "be" {
        fmt.Println(termenv.String("  Прадмет не знойдзены!").Foreground(termenv.ANSIRed))
    }
}

func ShowInventory(player *data.Player) {
    if len(player.Inventory.Items) == 0 {
        if data.Lang == "en" {
            fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
        } else if data.Lang == "ua" {
            fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
        } else if data.Lang == "be" {
            fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
        }
        return
    }
    
    if data.Lang == "en" {
        fmt.Println(termenv.String("  Your inventory:").Foreground(termenv.ANSIMagenta).Bold())
    } else if data.Lang == "ua" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    } else if data.Lang == "be" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    }
    
    for _, item := range player.Inventory.Items {
        fmt.Printf(" [%d] %s x%d\n", item.ID, item.Name, item.Quantity)
    }
}
