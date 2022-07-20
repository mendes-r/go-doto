package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Box Demo").
		SetBackgroundColor(tcell.Color100)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
