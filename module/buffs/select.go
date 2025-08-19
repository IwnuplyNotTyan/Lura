package buff

import (
	"fmt"

	"Lura/data"

	"github.com/charmbracelet/huh"
)

func selectBuff(player *data.Player) []string {
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

