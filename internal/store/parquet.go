package store

import (
	_ "embed"
	"fmt"

	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

//go:embed schema.json
var schema string

type Message struct {
	ID           string
	Title        string
	Lang         string
	RevisionDate string
	Desc         string
	Publisher    string
	Src          string
	Books        []Book
}

type Book struct {
	ID       string    `xml:"osisID,attr"`
	Chapters []Chapter `xml:"chapter"`
}

type Chapter struct {
	Text   string  `xml:",chardata"`
	OsisID string  `xml:"osisID,attr"`
	Verses []Verse `xml:"verse"`
}

type Verse struct {
	ID   string `xml:"osisID,attr"`
	Text string `xml:",chardata"`
}

func (m *Message) Write(src source.ParquetFile) error {
	w, err := writer.NewParquetWriter(src, schema, 1)
	if err != nil {
		return fmt.Errorf("schema invalid: %w", err)
	}
	w.RowGroupSize = 128 * 1024 * 1024
	w.CompressionType = parquet.CompressionCodec_SNAPPY

	if err = w.Write(m); err != nil {
		return err
	}

	if err = w.WriteStop(); err != nil {
		return err
	}

	return nil
}

func Read(src source.ParquetFile) (*Message, error) {
	r, err := reader.NewParquetReader(src, schema, 1)
	if err != nil {
		return nil, err
	}

	var x Message
	if err = r.Read(&x); err != nil {
		return &x, err
	}

	return &x, nil
}
