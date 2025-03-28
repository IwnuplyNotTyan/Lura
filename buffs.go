package main

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/huh"
	"github.com/muesli/termenv"
)

var (
	buff1 string
	buff2 string
	buff3 string
)

func getRandomBuff(player *Player, excludeBuffs ...string) string {
	var buffs []string
	if player.loc == 1 {
		if lang == "en" {
			buffs = []string{
				"Upgrade Weapon",
				"Longsword",
				"Crossbow",
				//"Random Weapon",
				"Broken heart",
				"Turtle scute",
			}
		} else if lang == "be" {
			buffs = []string{
				"Палепшыць зброю",
				//"Выпадковая зброя",
				"Разбітае сэрца",
				"Шчыт чарапахі",
			}
		} else {
			buffs = []string{
				"Покращити зброю",
				//"Випадкова зброя",
				"Розбите серце",
				"Щиток черепахи",
			}
		}
	} else if player.loc == 0 {
		if lang == "en" {
			buffs = []string{
				"Crystal heart",
				"Lotus",
				"Tears",
			}
		} else if lang == "be" {
			buffs = []string{
				"Кристалічна сэрца",
				"Лотас",
				"Слёзы",
			}
		} else if lang == "ua" {
			buffs = []string{
				"Кристалічне серце",
				"Лотос",
				"Сльози",
			}
		}
	}

	// Remove excluded buffs
	availableBuffs := make([]string, 0, len(buffs))
	for _, buff := range buffs {
		if !contains(excludeBuffs, buff) {
			availableBuffs = append(availableBuffs, buff)
		}
	}

	// If no buffs are available, return an empty string or handle as needed
	if len(availableBuffs) == 0 {
		return ""
	}

	// Return a random buff from available buffs
	return availableBuffs[rand.Intn(len(availableBuffs))]
}

func buffsAction(player *Player) {
	currentCoins(player)

	buff1 = getRandomBuff(player)
	buff2 = getRandomBuff(player)
	buff3 = getRandomBuff(player)

	selectedBuffs := selectBuff(player)
	if len(selectedBuffs) == 0 {
		noBuffDialog()
		return
	}

	for _, buff := range selectedBuffs {
		switch buff {

		case "Longsword":
			if player.Coins > 20 {
				getLongsword()
			} else {
				noBuffDialog()
			}

		case "Crossbow":
			if player.Coins > 20 {
				getCrossbow()
			} else {
				noBuffDialog()
			}

		case "Щиток черепахи", "Turtle scute", "Шчыт чарапахі":
			if player.Coins > 20 {
				player.Coins -= 20
				currentHp := player.HP
				player.HP += 50
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentHp, player.HP)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Crystal heart":
			if player.Coins > 50 {
				player.Coins -= 50
				player.heart = 2
				fmt.Println(termenv.String(fmt.Sprintf("  Your heart regenerate new power")))
			} else {
				noBuffDialog()
			}

		case "Lotus", "Лотус", "Лотас":
			if player.Coins > 10 {
				player.Coins -= 10
				currentMaxStamina := player.maxStamina
				player.maxStamina += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentMaxStamina, player.maxStamina)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Tears", "Сльози", "Слёзы":
			if player.Coins > 5 {
				player.Coins -= 5
				currentMaxHP := player.maxHP
				player.maxHP += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentMaxHP, player.maxHP)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Broken heart", "Розбите серце", "Разбітае сэрца":
			if player.Coins > 50 {
				player.heart = 0
				player.Coins -= 50
				fmt.Println(termenv.String(fmt.Sprintf("  No heart now.")))
			} else {
				noBuffDialog()
			}

		case "Upgrade Weapon", "Покращити зброю", "Палепшыць зброю":
			if player.Coins > 30 {
				player.Coins -= 30
				CurrentDamage := player.Damage
				player.Damage += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", CurrentDamage, player.Damage)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		default:
			noBuffDialog()
		}
	}
}

func selectBuff(player *Player) []string {
	var selectedBuffs []string

	buff1 := getRandomBuff(player)
	buff2 := getRandomBuff(player, buff1)
	buff3 := getRandomBuff(player, buff1, buff2)

	f := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(" Select card").
				Options(
					huh.NewOption(buff1, buff1),
					huh.NewOption(buff2, buff2),
					huh.NewOption(buff3, buff3),
				).
				Value(&selectedBuffs),
		),
	)

	if err := f.Run(); err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return selectedBuffs
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
