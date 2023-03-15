package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const url = "https://xkcd.com/%d/info.0.json"

func main() {
	lastPage := flag.Int("count", 10, "The number of page")
	flag.Parse()
	var (
		wg     sync.WaitGroup
		buf    bytes.Buffer
		client = http.DefaultClient
	)

	buf.WriteString("[")

	for i := 1; i <= *lastPage; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			currPage := fmt.Sprintf(url, i)
			res, err := client.Get(currPage)
			if err != nil {
				log.Printf("HTTP request error: %s\n", err)
				return
			}
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				log.Printf("Read response body error: %s\n", err)
				return
			}
			buf.WriteString(fmt.Sprintf("%s,", data))
			log.Printf("Write data from %q\n", currPage)
		}(i)
	}

	wg.Wait()
	buf.WriteString("]")
	os.WriteFile("out.json", buf.Bytes(), 0666)
}
