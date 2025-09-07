package dialog

import (
	"fmt"

	"Lura/data"

	"github.com/muesli/termenv"
)

func NoBuffDialog() {
	switch data.Lang {
		case "ru":
			fmt.Println(termenv.String("  Бафф не был применен.").Foreground(termenv.ANSIYellow))
		case "ua":
			fmt.Println(termenv.String("  Бафф не застосовано.").Foreground(termenv.ANSIYellow))
		case "be":
			fmt.Println(termenv.String("  Бафф не быў ужыты.").Foreground(termenv.ANSIYellow))
		default:
			fmt.Println(termenv.String("  No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func CurrentCoins(player *data.Player) {
	switch data.Lang {
		case "ru":
			fmt.Printf("  У вас %d монет\n", player.Coins)
		case "ua":
			fmt.Printf("  У тебе %d копiйок\n", player.Coins)
		case "be":
			fmt.Printf("  У вас %d манет\n", player.Coins)
		default:
			fmt.Printf("  You have %d coins\n", player.Coins)
	}
}

