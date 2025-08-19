package data

import (
	"fmt"

	"github.com/muesli/termenv"
	"github.com/charmbracelet/huh"
)

func (inv *Inventory) AddItem(name string, effect string, value int, price int) {
    // Check if item already exists
    for i, item := range inv.Items {
        if item.Name == name {
            inv.Items[i].Quantity++
            return
        }
    }
    
    // Add new item
    newItem := Item{
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

func (inv *Inventory) RemoveItem(id int, quantity int) bool {
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

func (player *Player) UseItem(id int) {
    for _, item := range player.Inventory.Items {
        if item.ID == id {
            switch item.Effect {
	    case "Material":
		    if Lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("  You get %s, +%d material!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if Lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("  Ви отримали %s, +%d матеріалів!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    } else if Lang == "be" {
			fmt.Println(termenv.String(fmt.Sprintf("  Вы атрымалі %s, +%d матэрыялаў!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
		    }
            case "heal":
                player.HP = min(player.HP+item.Value, player.MaxHP)
                if Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, healed %d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d HP!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "damage_boost":
                player.Damage += item.Value
                if Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, +%d damage for this fight!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d пошкоджень у цьому бою!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d пашкоджанняў у гэтым баі!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
                
            case "stamina_restore":
                player.Stamina = min(player.Stamina+item.Value, player.MaxStamina)
                if Lang == "en" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Used %s, restored %d stamina!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "ua" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Використано %s, +%d витривалостi!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                } else if Lang == "be" {
                    fmt.Println(termenv.String(fmt.Sprintf("  Выкарыстана %s, +%d вынослівасці!", item.Name, item.Value)).Foreground(termenv.ANSIGreen))
                }
            }
            
            player.Inventory.RemoveItem(id, 1)
            return
        }
    }
    
    if Lang == "en" {
        fmt.Println(termenv.String("  Item not found!").Foreground(termenv.ANSIRed))
    } else if Lang == "ua" {
        fmt.Println(termenv.String("  Предмет не знайдено!").Foreground(termenv.ANSIRed))
    } else if Lang == "be" {
        fmt.Println(termenv.String("  Прадмет не знойдзены!").Foreground(termenv.ANSIRed))
    }
}

func ShowInventory(player *Player) {
    if len(player.Inventory.Items) == 0 {
        if Lang == "en" {
            fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
        } else if Lang == "ua" {
            fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
        } else if Lang == "be" {
            fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
        }
        return
    }
    
    if Lang == "en" {
        fmt.Println(termenv.String("  Your inventory:").Foreground(termenv.ANSIMagenta).Bold())
    } else if Lang == "ua" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    } else if Lang == "be" {
        fmt.Println(termenv.String("  Ваш інвентар:").Foreground(termenv.ANSIMagenta).Bold())
    }
    
    for _, item := range player.Inventory.Items {
        fmt.Printf(" [%d] %s x%d\n", item.ID, item.Name, item.Quantity)
    }
}

func UseItemMenu(player *Player) {
    if len(player.Inventory.Items) == 0 {
        if Lang == "en" {
            fmt.Println(termenv.String("  Your inventory is empty").Foreground(termenv.ANSIYellow))
        } else if Lang == "ua" {
            fmt.Println(termenv.String("  Ваш інвентар порожній").Foreground(termenv.ANSIYellow))
        } else if Lang == "be" {
            fmt.Println(termenv.String("  Ваш інвентар пусты").Foreground(termenv.ANSIYellow))
        }
        return
    }
    
    var selectedID int
    options := make([]huh.Option[int], 0, len(player.Inventory.Items)+1)
    
    // Add "Cancel" option
    if Lang == "en" {
        options = append(options, huh.NewOption("Cancel", -1))
    } else if Lang == "ua" {
        options = append(options, huh.NewOption("Скасувати", -1))
    } else if Lang == "be" {
        options = append(options, huh.NewOption("Адмяніць", -1))
    }
    
    // Add item options
    for _, item := range player.Inventory.Items {
        var desc string
        if Lang == "en" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescription(item.Effect), item.Value)
        } else if Lang == "ua" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionUA(item.Effect), item.Value)
        } else if Lang == "be" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionBE(item.Effect), item.Value)
        }
        options = append(options, huh.NewOption(desc, item.ID))
    }
    
    var title string
    if Lang == "en" {
        title = " Select item to use"
    } else if Lang == "ua" {
        title = " Виберіть предмет для використання"
    } else if Lang == "be" {
        title = " Выберыце прадмет для выкарыстання"
    }
    
    f := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[int]().
                Title(title).
                Options(options...).
                Value(&selectedID),
        ),
    )
    
    if err := f.Run(); err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    if selectedID != -1 {
        player.UseItem(selectedID)
    }
}

func getEffectDescription(effect string) string {
    switch effect {
    case "heal":
        return "Heals"
    case "damage_boost":
        return "Boosts damage by"
    case "stamina_restore":
        return "Restores stamina by"
    default:
        return effect
    }
}

func getEffectDescriptionUA(effect string) string {
    switch effect {
    case "heal":
        return "Лікує"
    case "damage_boost":
        return "Збільшує пошкодження на"
    case "stamina_restore":
        return "Відновлює витривалість на"
    default:
        return effect
    }
}

func getEffectDescriptionBE(effect string) string {
    switch effect {
    case "heal":
        return "Лякуе"
    case "damage_boost":
        return "Павялічвае пашкоджанні на"
    case "stamina_restore":
        return "Аднаўляе вынослівасць на"
    default:
        return effect
    }
}
