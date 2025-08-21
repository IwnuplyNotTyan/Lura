package mods

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	lua "github.com/yuin/gopher-lua"
)

var loadedMods []string

func AutoLoadMods(L *lua.LState) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	modsDir := filepath.Join(homeDir, ".config", "lura", "mods")
	loadedMods = []string{}

	if err := os.MkdirAll(modsDir, 0755); err != nil {
		return fmt.Errorf("failed to create mods directory: %v", err)
	}

	err = filepath.Walk(modsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".lua" {
			return nil
		}

		if err := L.DoFile(path); err != nil {
			log.Printf("Failed to load mod %s: %v", path, err)
		} else {
			loadedMods = append(loadedMods, filepath.Base(path))
		}

		return nil
	})

	return err
}

func ModsLoaded() bool {
    return len(loadedMods) > 0
}

func GetLoadedMods() []string {
	return loadedMods
}
