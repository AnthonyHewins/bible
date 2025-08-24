package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/AnthonyHewins/bible/internal/codex"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func run() error {
	var wrapper osis
	if err := xml.NewDecoder(os.Stdin).Decode(&wrapper); err != nil {
		return err
	}

	t := wrapper.OsisText
	w := codex.NewWriter(os.Stdout)

	for _, v := range t.Books {
		chapters := make([][]string, len(v.Chapters))
		for i, v := range v.Chapters {
			verses := make([]string, len(v.Verses))
			for i, v := range v.Verses {
				verses[i] = v.Text
			}

			chapters[i] = verses
		}

		w.Write(codex.Book{ID: v.ID, Text: chapters})
	}

	return w.Close()
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
		Books []book `xml:"div"`
	} `xml:"osisText"`
}

type book struct {
	ID       string    `xml:"osisID,attr"`
	Chapters []chapter `xml:"chapter"`
}

type chapter struct {
	Text   string  `xml:",chardata"`
	OsisID string  `xml:"osisID,attr"`
	Verses []verse `xml:"verse"`
}

type verse struct {
	ID   string `xml:"osisID,attr"`
	Text string `xml:",chardata"`
}
