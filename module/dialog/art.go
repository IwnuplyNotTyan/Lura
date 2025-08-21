package dialog

import (
	"fmt"
	"os"

	"Lura/assets"

	"github.com/charmbracelet/log"
)

func CaveArt() {
	data, err := asset.FS.ReadFile("preview/cave.txt")
	if err != nil {
		log.Info("Error reading file:", err)
		return
	}

	fmt.Print(string(data))
	os.Stdout.Sync()
}

func ForestArt() {
	data, err := asset.FS.ReadFile("preview/forest.txt")
	if err != nil {
		log.Info("Error reading file:", err)
		return
	}
	fmt.Print(string(data))
	os.Stdout.Sync()
}

func CatArt() {
	data, err := asset.FS.ReadFile("preview/cat.txt")
	if err != nil {
		log.Info("Error reading file:", err)
		return
	}
	fmt.Print(string(data))
	os.Stdout.Sync()
}

