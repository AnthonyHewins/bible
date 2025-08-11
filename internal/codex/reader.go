package codex

import (
	"errors"
	"fmt"
	"io"

	"github.com/parquet-go/parquet-go"
)

type Reader struct {
	r *parquet.GenericReader[buffer]
}

func NewReader(r io.ReaderAt) *Reader {
	return &Reader{r: parquet.NewGenericReader[buffer](r)}
}

func (r *Reader) ReadAll() ([]Book, error) {
	books := []Book{}

	b, book, chapter, verses := make([]buffer, 1000), Book{ID: BookNameGen}, 1, []string{}
	var eof bool
	for _, err := r.r.Read(b); !eof; _, err = r.r.Read(b) {
		switch {
		case err == nil:
		case errors.Is(err, io.EOF):
			eof = true
		default:
			return nil, err
		}

		for _, line := range b {
			if line.BookID < uint32(book.ID) {
				return books, fmt.Errorf("translation not sorted: looking for book %s and found %s", book.ID, BookName(line.BookID))
			}

			if line.Chapter < uint32(chapter) {
				return books, fmt.Errorf("translation not sorted: looking for chapter %d in %s but got %d in %s", chapter, book.ID, line.Chapter, BookName(line.BookID))
			}

			if line.BookID > uint32(book.ID) {
				books = append(books, book)
				book, chapter = Book{ID: BookName(line.BookID)}, 1
				continue
			}

			if line.Chapter > uint32(chapter) {
				book.Text = append(book.Text, verses)
				chapter, verses = 1, []string{}
				continue
			}

			verses = append(verses, line.Verse)
		}
	}

	return books, nil
}
