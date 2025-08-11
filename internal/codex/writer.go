package codex

import (
	"io"

	"github.com/parquet-go/parquet-go"
)

type Writer struct {
	w *parquet.GenericWriter[buffer]
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: parquet.NewGenericWriter[buffer](w, parquet.BloomFilters(
		parquet.SplitBlockFilter(7, "id"),
		parquet.SplitBlockFilter(8, "chapter"),
	))}
}

func (w *Writer) Write(rows ...Book) error {
	x := []buffer{}
	for _, book := range rows {
		for j, v := range book.Text {
			for _, verse := range v {
				x = append(x, buffer{
					BookID:  uint32(book.ID),
					Chapter: uint32(j + 1),
					Verse:   verse,
				})
			}
		}
	}

	_, err := w.w.Write(x)
	return err
}

func (w *Writer) Close() error {
	return w.w.Close()
}
