package dialog

import (
	"fmt"

	"Lura/data"

	"github.com/muesli/termenv"
)

func NoBuffDialog() {
	if data.Lang == "ua" {
		fmt.Println(termenv.String("  Бафф не застосовано.").Foreground(termenv.ANSIYellow))
	} else if data.Lang == "en" {
		fmt.Println(termenv.String("  Бафф не был применен.").Foreground(termenv.ANSIYellow))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String("  Бафф не быў ужыты.").Foreground(termenv.ANSIYellow))
	} else {
		fmt.Println(termenv.String("  No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func CurrentCoins(player *data.Player) {
	if data.Lang == "ru" {
		fmt.Printf("  У вас %d монет\n", player.Coins)
	} else if data.Lang == "ua" {
		fmt.Printf("  У тебе %d копiйок\n", player.Coins)
	} else if data.Lang == "be" {
		fmt.Printf("  У вас %d манет\n", player.Coins)
	} else {
		fmt.Printf("  You have %d coins\n", player.Coins)
	}
}

