package fight

import (
	"fmt"

	"Lura/data"
	"Lura/module/rng"
	"Lura/module/dialog"

	"github.com/charmbracelet/huh"
)

func SelectAttack(player *data.Player) string {
	var selectedAttack string
	if data.Lang == "ua" {
		Attack = "Атакувати"
		Heal = "Лікуватися"
		Defend = "Захищатися"
		Skip = "Пропустити"
	} else if data.Lang == "be" {
		Attack = "Атакаваць"
		Heal = "Вылечвацца"
		Defend = "Абараняцца"
		Skip = "Прапусціць"
	} else if data.Lang == "ru" {
		Attack = "Атаковать"
		Heal = "Лечиться"
		Defend = "Защищаться"
		Skip = "Пропустить"
	} else {
		Attack = "Attack"
		Defend = "Defend"
		Heal = "Heal"
		Skip = "Skip"
	}
	
	var f *huh.Form
	if !player.Monster {
		f = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(" Select action").
					Options(
						huh.NewOption(Attack, Attack),
						huh.NewOption(Defend, Defend),
						huh.NewOption(Heal, Heal),
						huh.NewOption(Skip, Skip),
					).
					Value(&selectedAttack),
			),
		)
	} else {
		f = huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(" Select action").
					Options(
						huh.NewOption(Attack, Attack),
						huh.NewOption(Heal, Heal),
						huh.NewOption(Defend, Defend),
					).
					Value(&selectedAttack),
			),
		)
	}
	
	if err := f.Run(); err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	dialog.ClearScreen()
	return selectedAttack
}

func takeWeapon(player *data.Player, monster *data.Monster) {
	var confirm bool
	var a, b, c string

	switch {
	case data.Lang == "ua":
		a = "Ви хочете взяти зброю?"
		b = "Так"
		c = "Ні"
	case data.Lang == "ru":
		a = "Вы хотите взять оружие?"
		b = "Да"
		c = "Нет"
	case data.Lang == "be":
		a = "Вы хочаце ўзяць зброю?"
		b = "Так"
		c = "Не"
	default:
		a = "Do you want to take the weapon?"
		b = "Yes"
		c = "No"
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(a).
				Affirmative(b).
				Negative(c).
				Value(&confirm),
		),
	).Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirm {
		if monster.ID == 17 {
			rng.GetLanter(player)
		} else if monster.ID == 1 {
			rng.GetMusket(player)
		}
	}
}

