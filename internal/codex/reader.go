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

	b, book, chapter, verses := make([]buffer, 1000), Book{}, 1, []string{}
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
			if book.ID == "" {
				book.ID = line.BookID
			} else if book.ID != line.BookID {
				books = append(books, book)
				book, chapter, verses = Book{ID: line.BookID}, 1, []string{line.Verse}
				continue
			}

			switch {
			case line.Chapter < uint32(chapter):
				return books, fmt.Errorf(
					"translation not sorted: looking for chapter %d (or %d in sequence) in %s but got %d in %s (corrupted data)",
					chapter,
					chapter+1,
					book.ID,
					line.Chapter,
					line.BookID,
				)
			case line.Chapter > uint32(chapter):
				book.Text = append(book.Text, verses)
				chapter, verses = int(line.Chapter), []string{line.Verse}
			default:
				verses = append(verses, line.Verse)
			}
		}
	}

	return books, nil
}
