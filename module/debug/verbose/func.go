package v

import (
	"Lura/module/data"

	"github.com/charmbracelet/log"
)

func Verbose(msg string) {
	if data.Verbose {
		log.Info(msg)
	}
}
