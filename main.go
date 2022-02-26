package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
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
	fmt.Println("scrape starting...\nscraper stats:\n\t", c)
	c.OnHTML("*", func(element *colly.HTMLElement) {
		// element.ForEach("li", func(_ int, el *colly.HTMLElement) {
		// 	writer.Write([]string{el.ChildText("a:nth=child(1)")})
		// })
		// fmt.Println("\nscrape done :)\n")
		fmt.Println(element)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://dictionary.com/")
	fmt.Println("scraper stats:\n\t", c)
	fmt.Println("program ending...")
}
