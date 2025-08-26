package data

import (
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/huh"
)

func SelectLanguage() string {
	var selectedLang string
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(" Select Language").
				Options(
					huh.NewOption("English", "en"),
					huh.NewOption("Українська", "ua"),
					huh.NewOption("Беларускiй", "be"),
					huh.NewOption("Русский", "ru"),
				).
				Value(&selectedLang),
		),
	)

	if err := f.Run(); err != nil {
		log.Printf("Error selecting language: %v", err)
		return "en"
	}

	switch selectedLang {
	case "en", "ua", "be", "ru":
		return selectedLang
	default:
		return "en"
	}
}

