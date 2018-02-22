package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	nsDC      = "http://purl.org/dc/elements/1.1/"
	nsSynd    = "http://purl.org/rss/1.0/modules/syndication/"
	nsAdmin   = "http://webns.net/mvcb/"
	nsAtom    = "http://www.w3.org/2005/Atom"
	nsRDF     = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	nsContent = "http://purl.org/rss/1.0/modules/content/"
	nsITunes  = "http://www.itunes.com/dtds/podcast-1.0.dtd"
)

type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	URL     string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

type Description struct {
	XMLName xml.Name `xml:"description"`
	Content string   `xml:",cdata"`
}

type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  int32    `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type ITunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	URL     string   `xml:"href,attr"`
}

type ITunesCategory struct {
	XMLName       xml.Name          `xml:"itunes:category"`
	Text          string            `xml:"text,attr"`
	Subcategories []*ITunesCategory `xml:"omitempty"`
}

type ITunesDuration struct {
	XMLName  xml.Name `xml:"itunes:duration"`
	Duration time.Duration
}

func (d *ITunesDuration) String() string {
	return "abcd"
}

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	DC      string   `xml:"xmlns:dc,attr"`
	Synd    string   `xml:"xmlns:sy,attr"`
	Admin   string   `xml:"xmlns:admin,attr"`
	Atom    string   `xml:"xmlns:atom,attr"`
	RDF     string   `xml:"xmlns:rdf,attr"`
	Content string   `xml:"xmlns:content,attr"`
	ITunes  string   `xml:"xmlns:itunes,attr"`
	Podcast *Podcast
}

type Podcast struct {
	XMLName          xml.Name `xml:"channel"`
	Title            string   `xml:"title"`
	Copyright        string   `xml:"copyright"`
	Link             string   `xml:"link"`
	AtomLink         *AtomLink
	Language         string `xml:"language"`
	Description      string `xml:"description"`
	ITunesSubtitle   string `xml:"itunes:subtitle"`
	ITunesAuthor     string `xml:"itunes:author"`
	ITunesSummary    string `xml:"itunes:summary"`
	ITunesExplicit   bool   `xml:"itunes:explicit"`
	ITunesImage      *ITunesImage
	ITunesKeywords   string `xml:"itunes:keywords"`
	ITunesOwnerName  string `xml:"itunes:owner>itunes:name"`
	ITunesOwnerEmail string `xml:"itunes:owner>itunes:email"`
	ITunesCategory   *ITunesCategory
	Episodes         []*Episode
}

type Episode struct {
	XMLName        xml.Name  `xml:"item"`
	Title          string    `xml:"title"`
	Link           string    `xml:"link"`
	Guid           string    `xml:"guid"`
	PubDate        time.Time `xml:"pubDate"`
	Creator        string    `xml:"dc:creator"`
	Description    *Description
	Enclosure      *Enclosure
	ITunesSubtitle string `xml:"itunes:subtitle"`
	ITunesAuthor   string `xml:"itunes:author"`
	ITunesImage    *ITunesImage
	// ITunesDuration time.Duration `xml:"itunes:duration"` // TOOD: should be formatted as HH:MM (eg. 10:10)
	ITunesDuration *ITunesDuration
	ITunesExplicit bool   `xml:"itunes:explicit"`
	ITunesKeywords string `xml:"itunes:keywords"`
	ITunesSummary  string `xml:"itunes:summary"`
}

func main() {
	buf := bytes.NewBufferString(xml.Header)
	feed := &Feed{
		DC:      nsDC,
		Synd:    nsSynd,
		Admin:   nsAdmin,
		Atom:    nsAtom,
		RDF:     nsRDF,
		Content: nsContent,
		ITunes:  nsITunes,
		Podcast: &Podcast{
			Title:     "My Podcast",
			Copyright: "All rights reserved",
			Link:      "http://podcast.com",
			AtomLink: &AtomLink{
				URL:  "http://podcast.com/atom",
				Rel:  "alternate",
				Type: "application/rss+xml",
			},
			Language:         "en-us",
			Description:      "Description of my podcast",
			ITunesSubtitle:   "iTunes subtitle text",
			ITunesAuthor:     "iTunes author text",
			ITunesSummary:    "iTunes summary text",
			ITunesExplicit:   false,
			ITunesImage:      &ITunesImage{URL: "http://podcast.com/static/image.jpg"},
			ITunesKeywords:   strings.Join([]string{"keyword1", "keyword2"}, ","),
			ITunesOwnerName:  "me",
			ITunesOwnerEmail: "me@email.com",
			ITunesCategory: &ITunesCategory{
				Text: "category",
				Subcategories: []*ITunesCategory{
					&ITunesCategory{Text: "subcategory1"},
					&ITunesCategory{Text: "subcategory2"},
				},
			},
			Episodes: []*Episode{
				&Episode{
					Title:       "Episode 1 - First Episode",
					Link:        "https://podcast.com/episode-1",
					Guid:        "podcast.com/episodes/1",
					PubDate:     time.Now().UTC(),
					Creator:     "Me",
					Description: &Description{Content: "Description of an episode"},
					Enclosure: &Enclosure{
						URL:    "https://cdn.podcast.com/uploads/podcast/1/podcast-1.mp3",
						Length: 69374228,
						Type:   "audio/mpeg",
					},
					ITunesSubtitle: "iTunes subtitle text",
					ITunesAuthor:   "iTunes author text",
					ITunesImage:    &ITunesImage{URL: "http://podcast.com/static/podcast-1.jpg"},
					ITunesDuration: &ITunesDuration{Duration: time.Duration(77 * time.Minute)},
					ITunesExplicit: false,
					ITunesKeywords: strings.Join([]string{"keyword1", "keyword2"}, ","),
					ITunesSummary:  "iTunes summary text",
				},
				&Episode{
					Title:       "Episode 2 - Second Episode",
					Link:        "https://podcast.com/episode-2",
					Guid:        "podcast.com/episodes/2",
					PubDate:     time.Now().UTC(),
					Creator:     "Me",
					Description: &Description{Content: "Description of an episode"},
					Enclosure: &Enclosure{
						URL:    "https://cdn.podcast.com/uploads/podcast/2/podcast-2.mp3",
						Length: 69374228,
						Type:   "audio/mpeg",
					},
					ITunesSubtitle: "iTunes subtitle text",
					ITunesAuthor:   "iTunes author text",
					ITunesImage:    &ITunesImage{URL: "http://podcast.com/static/podcast-2.jpg"},
					ITunesDuration: &ITunesDuration{Duration: time.Duration(77 * time.Minute)},
					ITunesExplicit: false,
					ITunesKeywords: strings.Join([]string{"keyword1", "keyword2"}, ","),
					ITunesSummary:  "iTunes summary text",
				},
			},
		},
	}

	bs, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	buf.Write(bs)
	fmt.Println(buf.String())
}
