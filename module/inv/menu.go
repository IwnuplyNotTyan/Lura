package inv

import (
	"fmt"

	"Lura/"

	"github.com/muesli/termenv"
	"github.com/charmbracelet/huh"
)

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

