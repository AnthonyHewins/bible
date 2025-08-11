package main

import (
	"fmt"
	"os"

	"github.com/AnthonyHewins/bible/internal/codex"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	r, err := codex.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(r[0])
		panic(err)
	}
	fmt.Println(r[0].Text[0])
	// app := tview.NewApplication()

	// home := tview.NewBox().SetBorder(true).SetTitle("† Holy Bible †")

	// books := tview.NewList().ite

	// home.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	// 	switch event.Rune() {
	// 	// case 's':
	// 	case 'o':
	// 		pages.SwitchToPage()
	// 	case 'q':
	// 		app.Stop()
	// 	}

	// 	return event
	// })

	// pages := tview.NewPages().
	// 	AddAndSwitchToPage("Homepage", home, true)

	// // tview.NewInputField().SetLabel("Go to book: ").SetFieldWidth(30).Set

	// app.SetRoot(home, true)

	// if err := app.Run(); err != nil {
	// 	panic(err)
	// }
}
