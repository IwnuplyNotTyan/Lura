package dialog

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetLine(lines []string, index int) string {
	if index < len(lines) {
		return lines[index]
	}
	return ""
}

