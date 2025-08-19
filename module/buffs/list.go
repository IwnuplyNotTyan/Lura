package buff

import (
	"math/rand"

	"Lura/data"
)

var (
	buff1 string
	buff2 string
	buff3 string
)

func getRandomBuff(player *data.Player, excludeBuffs ...string) string {
	var buffs []string
	if player.Loc == 0 {
		if data.Lang == "en" {
			buffs = []string{
				"Upgrade Weapon",
				"Longsword",
				"Crossbow",
				//"Random Weapon",
				"Broken heart",
				"Turtle scute",
			}
		} else if data.Lang == "be" {
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
				"Довгий меч", "Арбалет",
			}
		}
	} else if player.Loc == 1 {
		if data.Lang == "en" {
			buffs = []string{
				"Crystal heart",
				"Lotus",
				"Tears",
				//"Amethyst necklace",
				//"Flask with star tears",
			}
		} else if data.Lang == "be" {
			buffs = []string{
				"Кристалічнае сэрца",
				"Лотас",
				"Слёзы",
			}
		} else if data.Lang == "ua" {
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

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
