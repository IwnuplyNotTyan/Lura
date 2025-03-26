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

func getRandomBuff(player *Player) string {
	var buffs []string
	if player.loc == 1 {
		if lang == "en" {
			buffs = []string{
				"Upgrade Weapon",
				"Random Weapon",
				"Broken heart",
				//"Pearl necklace",
				"Turtle scute",
			}
		} else if lang == "be" {
			buffs = []string{
				"Палепшыць зброю",
				"Выпадковая зброя",
				"Разбітае сэрца",
				"Шчыт чарапахі",
			}
		} else {
			buffs = []string{
				"Покращити зброю",
				"Випадкова зброя",
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
	return buffs[rand.Intn(len(buffs))]
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
		case "Random Weapon", "Випадкова зброя", "Выпадковая зброя":
			if player.Coins > 10 {
				player.Coins -= 10
				currentWeapon := player.WeaponType
				weaponType, weaponDamage := getRandomWeapon()
				player.WeaponType = weaponType
				player.Damage = weaponDamage
				fmt.Println(termenv.String(fmt.Sprintf("  %s  %s", currentWeapon, player.WeaponType)).Foreground(termenv.ANSIGreen))
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
	buff2 := getRandomBuff(player)
	buff3 := getRandomBuff(player)

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
