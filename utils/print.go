package utils

import (
	"github.com/pterm/pterm"
)

func SuperPrint(super interface{}) {
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Margin:          20,
	}
	newHeader.Println(super)
}
