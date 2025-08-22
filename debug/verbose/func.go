package v

import (
	"Lura/data"

	"github.com/charmbracelet/log"
)

func Verbose(msg string) {
	if data.Verbose {
		log.Info(msg)
	}
}
