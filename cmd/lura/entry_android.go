//go:build android

package main

import (
    "Lura/lurarun"
    "Lura/termbridge"
)

// On Android, register the CLI's main as the bridge entrypoint, so the Android app
// can start the game without modifying game code.
func init() {
    termbridge.RegisterEntrypoint(func() {
        lurarun.Run()
    })
}

