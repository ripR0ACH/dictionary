package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	alph := [27]string{"0", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	data := "data.csv"
	file, err := os.Create(data)
	if err != nil {
		log.Fatalf("Could not create file, \nerr: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	fmt.Println("scrape starting...\n")
	c.OnHTML("main ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{el.ChildText("a:first-child")})
		})
		fmt.Println("\nscrape done :)\n")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	scraping := true
	c.OnError(func(r *colly.Response, err error) {
		print(r)
		scraping = false
	})
	for i := 0; i < len(alph); i++ {
		ind := 1
		for scraping {
			c.Visit("https://dictionary.com/list/" + alph[i] + "/" + strconv.Itoa(ind))
			ind += 1
		}
		scraping = true
	}
	fmt.Println("program ending...")
}
