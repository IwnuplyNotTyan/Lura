package fight

import (
	"Lura/module/data"
	"Lura/module/rng"
)

func enemyTurn(monster *data.Monster, player *data.Player) string {
	if monster.ID == 10 {
		rngChoice := rng.Rng() % 3
		switch rngChoice {
		case 0: return "Attack"
		case 1: return "Defend"
		case 2: return "Heal"
		default: return "Attack"
	}
	} else {
		if player.Stamina < 10 {
			switch rng.Rng2() {
				case 1: return "Attack"
				case 2: return "Heal"
				default: return "Defend"}
		} else if monster.HP > 15 || player.HP < 15 {
			switch rng.Rng2() {
				case 1: return "Attack"
				case 2: return "Defend"
				default: return "Heal"}
		} else if monster.Position == player.Position-1 {
			switch rng.Rng2() {
				case 1: return "Defend"
				case 2: return "Attack"
				default: return "Heal"}
		} else if monster.HP < 15 {
			switch rng.Rng2() {
				case 1: return "Heal"
				case 2: return "Defend"
				default: return "Attack"}
		} else {
			return "Attack"
		}
	}
}

