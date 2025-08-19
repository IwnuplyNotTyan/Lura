package buff

import (
	"fmt"

	"Lura/data"

	"github.com/muesli/termenv"
)

func noBuffDialog() {
	if data.Lang == "ua" {
		fmt.Println(termenv.String("  Бафф не застосовано.").Foreground(termenv.ANSIYellow))
	} else if data.Lang == "en" {
		fmt.Println(termenv.String("  No Buff Applied.").Foreground(termenv.ANSIYellow))
	} else if data.Lang == "be" {
		fmt.Println(termenv.String("  Бафф не быў ужыты.").Foreground(termenv.ANSIYellow))
	}
}

func currentCoins(player *data.Player) {
	if data.Lang == "en" {
		fmt.Printf("  You have %d coins\n", player.Coins)
	} else if data.Lang == "ua" {
		fmt.Printf("  У тебе %d копiйок\n", player.Coins)
	} else if data.Lang == "be" {
		fmt.Printf("  У вас %d манет\n", player.Coins)
	}
}

