package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/AnthonyHewins/bible/gen/go/bible/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	books := make([]*bible.Book, len(t.Books))
	for i, v := range t.Books {
		id := bookToID(v.ID)
		if id == -1 {
			return fmt.Errorf("invalid bible book ID: '%s'. Should be chapter abbreviation like Gen, Acts, etc", v.ID)
		}

		chapters := make([]*bible.Chapter, len(v.Chapters))
		for i, v := range v.Chapters {
			verses := make([]string, len(v.Verses))
			for i, v := range v.Verses {
				verses[i] = v.Text
			}

			chapters[i] = &bible.Chapter{Verses: verses}
		}

		books[i] = &bible.Book{Id: id, Chapters: chapters}
	}

	d, err := time.Parse(time.DateOnly, t.Header.RevisionDesc.Date)
	if err != nil {
		return err
	}

	m := bible.Translation{
		Id:           t.OsisIDWork,
		Title:        t.Header.Work.Title,
		Lang:         t.Lang,
		RevisionDate: timestamppb.New(d),
		Desc:         t.Header.Work.Description,
		Publisher:    t.Header.Work.Publisher,
		Src:          t.Header.Work.Source,
		Books:        books,
	}

	buf, err := proto.Marshal(&m)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(buf)
	return err
}

func bookToID(s string) bible.BookName {
	switch s {
	case "Gen":
		return bible.BookName_BOOK_NAME_GENESIS
	case "Exod":
		return bible.BookName_BOOK_NAME_EXODUS
	case "Lev":
		return bible.BookName_BOOK_NAME_LEVITICUS
	case "Num":
		return bible.BookName_BOOK_NAME_NUMBERS
	case "Deut":
		return bible.BookName_BOOK_NAME_DEUTERONOMY
	case "Josh":
		return bible.BookName_BOOK_NAME_JOSHUA
	case "Judg":
		return bible.BookName_BOOK_NAME_JUDGES
	case "Ruth":
		return bible.BookName_BOOK_NAME_RUTH
	case "1Sam":
		return bible.BookName_BOOK_NAME_1_SAMUEL
	case "2Sam":
		return bible.BookName_BOOK_NAME_2_SAMUEL
	case "1Kgs":
		return bible.BookName_BOOK_NAME_1_KINGS
	case "2Kgs":
		return bible.BookName_BOOK_NAME_2_KINGS
	case "1Chr":
		return bible.BookName_BOOK_NAME_1_CHRONICLES
	case "2Chr":
		return bible.BookName_BOOK_NAME_2_CHRONICLES
	case "Ezra":
		return bible.BookName_BOOK_NAME_EZRA
	case "Neh":
		return bible.BookName_BOOK_NAME_NEHEMIAH
	case "Esth":
		return bible.BookName_BOOK_NAME_ESTHER
	case "Job":
		return bible.BookName_BOOK_NAME_JOB
	case "Ps":
		return bible.BookName_BOOK_NAME_PSALM
	case "Prov":
		return bible.BookName_BOOK_NAME_PROVERBS
	case "Eccl":
		return bible.BookName_BOOK_NAME_ECCLESIASTES
	case "Song":
		return bible.BookName_BOOK_NAME_SONG_OF_SONGS
	case "Isa":
		return bible.BookName_BOOK_NAME_ISAIAH
	case "Jer":
		return bible.BookName_BOOK_NAME_JEREMIAH
	case "Lam":
		return bible.BookName_BOOK_NAME_LAMENTATIONS
	case "Ezek":
		return bible.BookName_BOOK_NAME_EZEKIEL
	case "Dan":
		return bible.BookName_BOOK_NAME_DANIEL
	case "Hos":
		return bible.BookName_BOOK_NAME_HOSEA
	case "Joel":
		return bible.BookName_BOOK_NAME_JOEL
	case "Amos":
		return bible.BookName_BOOK_NAME_AMOS
	case "Obad":
		return bible.BookName_BOOK_NAME_OBADIAH
	case "Jonah":
		return bible.BookName_BOOK_NAME_JONAH
	case "Mic":
		return bible.BookName_BOOK_NAME_MICAH
	case "Nah":
		return bible.BookName_BOOK_NAME_NEHEMIAH
	case "Hab":
		return bible.BookName_BOOK_NAME_HABAKKUK
	case "Zeph":
		return bible.BookName_BOOK_NAME_ZEPHANIAH
	case "Hag":
		return bible.BookName_BOOK_NAME_HAGGAI
	case "Zech":
		return bible.BookName_BOOK_NAME_ZECHARIAH
	case "Mal":
		return bible.BookName_BOOK_NAME_MALACHI
	case "Matt":
		return bible.BookName_BOOK_NAME_MATTHEW
	case "Mark":
		return bible.BookName_BOOK_NAME_MARK
	case "Luke":
		return bible.BookName_BOOK_NAME_LUKE
	case "John":
		return bible.BookName_BOOK_NAME_JOHN
	case "Acts":
		return bible.BookName_BOOK_NAME_ACTS
	case "Rom":
		return bible.BookName_BOOK_NAME_ROMANS
	case "1Cor":
		return bible.BookName_BOOK_NAME_1_CORINTHIANS
	case "2Cor":
		return bible.BookName_BOOK_NAME_2_CORINTHIANS
	case "Gal":
		return bible.BookName_BOOK_NAME_GALATIANS
	case "Eph":
		return bible.BookName_BOOK_NAME_EPHESIANS
	case "Phil":
		return bible.BookName_BOOK_NAME_PHILIPPIANS
	case "Col":
		return bible.BookName_BOOK_NAME_COLOSSIANS
	case "1Thess":
		return bible.BookName_BOOK_NAME_1_THESSALONIANS
	case "2Thess":
		return bible.BookName_BOOK_NAME_2_THESSALONIANS
	case "1Tim":
		return bible.BookName_BOOK_NAME_1_TIMOTHY
	case "2Tim":
		return bible.BookName_BOOK_NAME_2_TIMOTHY
	case "Titus":
		return bible.BookName_BOOK_NAME_TITUS
	case "Phlm":
		return bible.BookName_BOOK_NAME_PHILEMON
	case "Heb":
		return bible.BookName_BOOK_NAME_HEBREWS
	case "Jas":
		return bible.BookName_BOOK_NAME_JAMES
	case "1Pet":
		return bible.BookName_BOOK_NAME_1_PETER
	case "2Pet":
		return bible.BookName_BOOK_NAME_2_PETER
	case "1John":
		return bible.BookName_BOOK_NAME_1_JOHN
	case "2John":
		return bible.BookName_BOOK_NAME_2_JOHN
	case "3John":
		return bible.BookName_BOOK_NAME_3_JOHN
	case "Jude":
		return bible.BookName_BOOK_NAME_JUDE
	case "Rev":
		return bible.BookName_BOOK_NAME_REVELATION
	default:
		return -1
	}
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
