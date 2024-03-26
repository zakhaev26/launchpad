package internal

import (
	"os"

	"github.com/charmbracelet/log"
)

func HandleError(err error, msg string) {
	if err != nil {
		log.Error(msg, err)
		os.Exit(1)
	}
}
