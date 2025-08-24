package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnthonyHewins/bible/internal/codex"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type reader struct {
	books               []codex.Book
	bookIdx, chapterIdx int
}

func (r *reader) title() string {
	b := r.books[r.bookIdx]
	return fmt.Sprintf("%s %d", b.ID, r.chapterIdx+1)
}

func (r *reader) prev() []string {
	if r.chapterIdx-1 >= 0 {
		r.chapterIdx--
		return r.books[r.bookIdx].Text[r.chapterIdx]
	}

	if r.bookIdx-1 >= 0 {
		r.bookIdx--
		r.chapterIdx = len(r.books[r.bookIdx].Text) - 1
	}

	return r.books[r.bookIdx].Text[r.chapterIdx]
}

func (r *reader) next() []string {
	chapters := len(r.books[r.bookIdx].Text)
	if r.chapterIdx+1 < chapters {
		r.chapterIdx++
		return r.books[r.bookIdx].Text[r.chapterIdx]
	}

	bookCount := len(r.books)
	if r.bookIdx+1 < bookCount {
		r.chapterIdx, r.bookIdx = 0, r.bookIdx+1
	}

	return r.books[r.bookIdx].Text[r.chapterIdx]
}

func run() error {
	books, err := readTranslation()
	if err != nil {
		return err
	}

	r := &reader{books: books}

	app := tview.NewApplication()
	// home := tview.NewBox().SetBorder(true).SetTitle("† Holy Bible †")

	gen := books[0]
	txtView := tview.NewTextView().SetText(strings.Join(gen.Text[0], "\n")).SetWordWrap(true)
	txtView.SetBorder(true).SetTitle(r.title())

	txtView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			txtView.SetText(strings.Join(r.next(), "\n")).SetTitle(r.title())
		case tcell.KeyLeft:
			txtView.SetText(strings.Join(r.prev(), "\n")).SetTitle(r.title())
		case tcell.KeyRune:
			switch event.Rune() {
			case 'b':
			// 	txtView = txtView.SetText(books[0].Text[0][0])
			// case 'o':
			// 	pages.SwitchToPage()
			case 'q':
				app.Stop()
			}
		}

		return event
	})

	app.SetRoot(txtView, true).EnableMouse(true)
	return app.Run()
}

func readTranslation() ([]codex.Book, error) {
	n := len(os.Args)
	if n < 2 {
		return nil, fmt.Errorf("Missing parquet file")
	}

	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	books, err := codex.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	if len(books) == 0 || len(books[0].Text) == 0 || len(books[0].Text[0]) == 0 {
		return nil, fmt.Errorf("data corrupted: first chapter of genesis missing")
	}

	return books, nil
}
