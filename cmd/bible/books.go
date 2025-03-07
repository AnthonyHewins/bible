package main

import "github.com/rivo/tview"

func bookList() *tview.List {
	list := tview.NewList()

	list.AddItem("", "", 0, func() {

	})

	return list
}
