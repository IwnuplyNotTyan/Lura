package lurarun

import (
    "flag"
    "fmt"

    "Lura/data"
    "Lura/fight"
    "Lura/module/debug"
    "Lura/module/mods"
    "Lura/module/rng"

    "github.com/charmbracelet/log"
    "github.com/charmbracelet/huh"
    lua "github.com/yuin/gopher-lua"
)

// Run mirrors the behavior of cmd/lura/main.go main() so that mobile builds can
// start the game loop without importing a main package.
func Run() {
    debugMode := flag.Bool("debug", false, "Enable debug shell")
    flag.Parse()

    L := lua.NewState()
    defer L.Close()

    // reset/clear screen
    // dialog.ClearScreen() is called in main, but to avoid importing dialog here
    // we replicate ClearScreen via ANSI escape.
    fmt.Print("\x1b[2J\x1b[H")

    player := data.Player{}
    cfg := data.TouchConfig(&player)

    data.Lang = cfg.Language
    if data.Lang == "" {
        fmt.Println()
        data.Lang = selectLanguage()
        cfg.Language = data.Lang
        data.SaveConfig(data.GetConfigPath(), cfg)
    }

    data.SeedData()
    mods.RegisterTypes(L)
    L.SetGlobal("lang", lua.LString(data.Lang))

    if err := mods.AutoLoadMods(L); err != nil {
        log.Fatalf("Failed to auto-load mods: %v", err)
    }

    weaponType, weaponDamage, weaponID := rng.GetRandomWeapon()
    hp := rng.RngHp()
    st := rng.RngHp()

    player = data.Player{
        WeaponType: weaponType,
        Damage:     weaponDamage * rng.Rng(),
        HP:         hp,
        MaxHP:      hp,
        Coins:      0,
        Stamina:    st,
        MaxStamina: st,
        Heart:      1,
        Loc:        0,
        WeaponID:   weaponID,
        Monster:    false,
        Position:   0,
        Inventory: data.Inventory{
            Items:  make([]data.Item, 0),
            NextID: 1,
        },
    }
    data.Tmp = -1

    if *debugMode {
        debug.DebugShell(L, &player)
    }
    fight.Fight(&player, nil, &data.Config{}, &data.Weapon{})
}

func selectLanguage() string {
    var selectedLang string
    f := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title(" Select Language").
                Options(
                    huh.NewOption("English", "en"),
                    huh.NewOption("Українська", "ua"),
                    huh.NewOption("Беларускiй", "be"),
                ).
                Value(&selectedLang),
        ),
    )

    if err := f.Run(); err != nil {
        log.Printf("Error selecting language: %v", err)
        return "en"
    }

    switch selectedLang {
    case "en", "ua", "be":
        return selectedLang
    default:
        return "en"
    }
}

