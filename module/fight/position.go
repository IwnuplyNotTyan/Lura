package fight

import (
	"fmt"
	"strings"

	"Lura/module/data"
)

func DisplayPositions(player *data.Player, monster *data.Monster) {
    positions := make([]string, 6)
    for i := range positions {
        positions[i] = " "
    }
    
    positions[player.Position] = " "
    positions[monster.Position] = " "
    
    fmt.Println(strings.Join(positions, ""))
}

