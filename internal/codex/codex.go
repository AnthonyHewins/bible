package codex

import (
	"time"
)

type Translation struct {
	ID           string
	Title        string
	Lang         string
	RevisionDate time.Time
	Desc         string
	Publisher    string
	Src          string
	Books        []Book
}

type Book struct {
	ID   string
	Text [][]string
}

type buffer struct {
	BookID  string `parquet:"id"`
	Chapter uint32 `parquet:"chapter"`
	Verse   string `parquet:"verse"`
}
