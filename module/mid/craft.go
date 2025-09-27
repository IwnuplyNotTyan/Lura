package mid

import (
	"Lura/module/data"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

func Crafting(player *data.Player) {
    var selections []string 
    err := huh.NewMultiSelect[string]().
        Title("Crafting").
        Options(
            huh.NewOption("Bell", "bell"),
        ).
        Value(&selections).
        Run()
    if err != nil {
        log.Info(err)
        return
    }
    
    for _, selection := range selections {
        if selection == "bell" {
            log.Info("Crafting a bell")
        }
    }
}

