package buff

import (
	"fmt"

	"Lura/data"
	"Lura/module/dialog"
	"Lura/module/rng"

	"github.com/muesli/termenv"
)

func BuffsAction(player *data.Player) {
	dialog.CurrentCoins(player)

	buff1 = getRandomBuff(player)
	buff2 = getRandomBuff(player)
	buff3 = getRandomBuff(player)

	selectedBuffs := selectBuff(player)
	if len(selectedBuffs) == 0 {
		dialog.NoBuffDialog()
		return
	}

	for _, buff := range selectedBuffs {
		switch buff {

		//case "Amethyst necklace":
		//	if player.Coins > 20 {
		//		player.amenuck = true
		//	} else {
		//		noBuffDialog()
		//	}

		//case "Flask with star tears":
		//	if player.Coins > 100 {
		//		maxh := player.maxHP
		//		d := player.Damage
		//		player.Damage *= 2
		//		player.maxHP += 2
		//		player.monster = true
		//		fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", maxh, player.maxHP)).Foreground(termenv.ANSIGreen))
		//		fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", d, player.Damage)).Foreground(termenv.ANSIGreen))
		//	}

		case "Longsword", "Довгий меч", "Доўгі меч":
			if player.Coins > 20 {
				w := player.WeaponType
				rng.GetLongsword(player)
				fmt.Println(termenv.String(fmt.Sprintf("󰓥  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Crossbow", "Арбалет":
			if player.Coins > 20 {
				w := player.WeaponType
				rng.GetCrossbow(player)
				fmt.Println(termenv.String(fmt.Sprintf("󱡁  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Щиток черепахи", "Turtle scute", "Шчыт чарапахі":
			if player.Coins > 20 {
				player.Coins -= 20
				currentHp := player.HP
				player.HP += 50
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentHp, player.HP)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Crystal heart", "Кристалічнае сэрца", "Кристалічне серце":
			if player.Coins > 50 {
				player.Coins -= 50
				player.Heart = 2
				fmt.Println(termenv.String("󰩖  Your heart regenerate new power").Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Lotus", "Лотус", "Лотас":
			if player.Coins > 10 {
				player.Coins -= 10
				currentMaxStamina := player.MaxStamina
				player.MaxStamina += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentMaxStamina, player.MaxStamina)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Tears", "Сльози", "Слёзы":
			if player.Coins > 5 {
				player.Coins -= 5
				currentMaxHP := player.MaxHP
				player.MaxHP += 10
				fmt.Println(termenv.String(fmt.Sprintf("󱐮  %d  %d", currentMaxHP, player.MaxHP)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Broken heart", "Розбите серце", "Разбітае сэрца":
			if player.Coins > 50 {
				player.Heart = 0
				player.Coins -= 50
				fmt.Println(termenv.String("  heart = false").Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		case "Upgrade Weapon", "Покращити зброю", "Палепшыць зброю":
			if player.Coins > 30 {
				player.Coins -= 30
				CurrentDamage := player.Damage
				player.Damage += 10
				fmt.Println(termenv.String(fmt.Sprintf("󰞇  %d  %d", CurrentDamage, player.Damage)).Foreground(termenv.ANSIGreen))
			} else {
				dialog.NoBuffDialog()
			}

		default:
			dialog.NoBuffDialog()
		}
	}
}

