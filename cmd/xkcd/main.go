package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// Comics represents xkcd comic strip
type Comics struct {
	ID         int `json:"num"`
	Day        string
	Month      string
	Year       string
	Link       string `json:",omitempty"`
	News       string `json:",omitempty"`
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
}

// String returns formated string that represent Comics
func (c *Comics) String() string {
	format := "Number - %d\nDate - %s/%s/%s\nLink - %s\nNews - %s\nTitle - %s\nSafe Title - %s\nTranscript - %s\nAlt - %s\nImg URL - %s\n"
	return fmt.Sprintf(format,
		c.ID,
		c.Month,
		c.Day,
		c.Year,
		c.Link,
		c.News,
		c.Title,
		c.SafeTitle,
		c.Transcript,
		c.Alt,
		c.Img,
	)
}

// ByNum represent a slice of Comics that uses for sorting
type ByNum []Comics

func (a ByNum) Len() int           { return len(a) }
func (a ByNum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNum) Less(i, j int) bool { return a[i].ID < a[j].ID }

func search(comics []Comics, id int) (Comics, error) {
	mid := len(comics) / 2
	switch {
	case len(comics) == 0:
		return Comics{}, errors.New("can not find comics in array. the length is 0")
	case comics[mid].ID > id:
		return search(comics[:mid], id)
	case comics[mid].ID < id:
		return search(comics[mid:], id)
	default:
		return comics[mid], nil
	}
}

func main() {
	page := flag.Int("p", 1, "The number of page")
	datapath := flag.String("f", "out.json", "The path to file with xkcd comics")
	flag.Parse()

	data, err := os.ReadFile(*datapath)
	if err != nil {
		log.Fatalf("Read file error: %s", err)
	}

	var comics []Comics
	if err := json.Unmarshal(data, &comics); err != nil {
		log.Fatalf("JSON Unmarshal error: %s", err)
	}

	sort.Sort(ByNum(comics))
	c, err := search(comics, *page)
	if err != nil {
		log.Fatalf("Search comics error: %s", err)
	}
	fmt.Println(c.String())
}
