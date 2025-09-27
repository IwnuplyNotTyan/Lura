package inv

import (
	"fmt"

	"Lura/module/data"

	"github.com/muesli/termenv"
	"github.com/charmbracelet/huh"
)

func UseItemMenu(player *data.Player) {
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
    
    var selectedID int
    options := make([]huh.Option[int], 0, len(player.Inventory.Items)+1)
    
    // Add "Cancel" option
    if data.Lang == "en" {
        options = append(options, huh.NewOption("Cancel", -1))
    } else if data.Lang == "ua" {
        options = append(options, huh.NewOption("Скасувати", -1))
    } else if data.Lang == "be" {
        options = append(options, huh.NewOption("Адмяніць", -1))
    }
    
    // Add item options
    for _, item := range player.Inventory.Items {
        var desc string
        if data.Lang == "en" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescription(item.Effect), item.Value)
        } else if data.Lang == "ua" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionUA(item.Effect), item.Value)
        } else if data.Lang == "be" {
            desc = fmt.Sprintf("%s (x%d) - %s %d", item.Name, item.Quantity, getEffectDescriptionBE(item.Effect), item.Value)
        }
        options = append(options, huh.NewOption(desc, item.ID))
    }
    
    var title string
    if data.Lang == "en" {
        title = " Select item to use"
    } else if data.Lang == "ua" {
        title = " Виберіть предмет для використання"
    } else if data.Lang == "be" {
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
        UseItem(player, selectedID)
    }
}

