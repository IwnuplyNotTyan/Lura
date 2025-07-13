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
	if player.loc == 0 {
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
				"Доўгі меч",
				"Арбалет",
			}
		} else {
			buffs = []string{
				"Покращити зброю",
				//"Випадкова зброя",
				"Розбите серце",
				"Щиток черепахи",
				"Довгий меч",
				"Арбалет",
			}
		}
	} else if player.loc == 1 {
		if lang == "en" {
			buffs = []string{
				"Crystal heart",
				"Lotus",
				"Tears",
				//"Amethyst necklace",
				//"Flask with star tears",
			}
		} else if lang == "be" {
			buffs = []string{
				"Кристалічнае сэрца",
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

	availableBuffs := make([]string, 0, len(buffs))
	for _, buff := range buffs {
		if !contains(excludeBuffs, buff) {
			availableBuffs = append(availableBuffs, buff)
		}
	}

	if len(availableBuffs) == 0 {
		return ""
	}

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
				getLongsword(player)
				fmt.Println(termenv.String(fmt.Sprintf("  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Crossbow", "Арбалет":
			if player.Coins > 20 {
				w := player.WeaponType
				getCrossbow(player)
				fmt.Println(termenv.String(fmt.Sprintf("  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
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

		case "Crystal heart", "Кристалічнае сэрца", "Кристалічне серце":
			if player.Coins > 50 {
				player.Coins -= 50
				player.heart = 2
				fmt.Println(termenv.String("  Your heart regenerate new power"))
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
				fmt.Println(termenv.String("  heart = false"))
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
