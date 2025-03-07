package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/AnthonyHewins/bible/internal/store"
	"github.com/xitongsys/parquet-go-source/local"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var wrapper osis
	if err := xml.NewDecoder(os.Stdin).Decode(&wrapper); err != nil {
		return err
	}

	t := wrapper.OsisText
	m := store.Message{
		ID:           t.OsisIDWork,
		Title:        t.Header.Work.Title,
		Lang:         t.Lang,
		RevisionDate: t.Header.RevisionDesc.Date,
		Desc:         t.Header.Work.Description,
		Publisher:    t.Header.Work.Publisher,
		Src:          t.Header.Work.Source,
		Books:        t.Books,
	}

	file, err := local.NewLocalFileWriter(os.Args[1])
	if err != nil {
		return err
	}

	if err = m.Write(file); err != nil {
		return err
	}

	f, err := local.NewLocalFileReader(os.Args[1])
	if err != nil {
		return err
	}

	x, err := store.Read(f)
	if err != nil {
		return err
	}

	fmt.Println(x.Books[0])
	return nil
}

type osis struct {
	XMLName        xml.Name `xml:"osis"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	OsisText       struct {
		Text        string `xml:",chardata"`
		OsisIDWork  string `xml:"osisIDWork,attr"`
		Lang        string `xml:"lang,attr"`
		OsisRefWork string `xml:"osisRefWork,attr"`
		Header      struct {
			Text         string `xml:",chardata"`
			RevisionDesc struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				P    string `xml:"p"`
			} `xml:"revisionDesc"`
			Work struct {
				Text        string `xml:",chardata"`
				OsisWork    string `xml:"osisWork,attr"`
				Title       string `xml:"title"`
				Contributor string `xml:"contributor"`
				Creator     []struct {
					Text string `xml:",chardata"`
					Role string `xml:"role,attr"`
				} `xml:"creator"`
				Subject     string `xml:"subject"`
				Date        string `xml:"date"`
				Description string `xml:"description"`
				Publisher   string `xml:"publisher"`
				Type        struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"type"`
				Identifier struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"identifier"`
				Source   string `xml:"source"`
				Language struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"language"`
				Relation  string `xml:"relation"`
				Coverage  string `xml:"coverage"`
				Rights    string `xml:"rights"`
				Scope     string `xml:"scope"`
				RefSystem string `xml:"refSystem"`
			} `xml:"work"`
		} `xml:"header"`
		Books []store.Book `xml:"div"`
	} `xml:"osisText"`
}
