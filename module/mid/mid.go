package mid

import (
	"Lura/data"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/huh"
)

func AfterLoc(player *data.Player) {
var confirm string

err := huh.NewSelect[string]().
	Title("Good night! What want?").
	Options(
		huh.NewOption("Sleep", "sleep"),
		huh.NewOption("Cooking", "cook"),
		huh.NewOption("Crafting", "craft"),
		huh.NewOption("Mining", "mine"),
	).
	Value(&confirm).
	Run()

if err != nil {
	log.Info(err)
	return
}

if confirm == "sleep" {
	log.Info(player.Stamina)
	staminaSleep := player.MaxStamina + 20
	player.Stamina = staminaSleep
	log.Info(player.Stamina)
} else if confirm == "craft" {
	Crafting(player)
} else if confirm == "mine" {
	log.Info("Mining...")
} else if confirm == "cook" {
	log.Info("Cooking...")
} else {
	log.Info("Invalid selection")
}
}

