package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	// alph := [27]string{"0", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	data := "data.csv"
	file, err := os.Create(data)
	if err != nil {
		log.Fatalf("Could not create file, \nerr: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	colly.IgnoreRobotsTxt()
	// word_collector := colly.NewCollector()
	def_collector := colly.NewCollector()
	def_collector.OnHTML("div.css-1avshm7.e16867sm0:nth-child(1)", func(e *colly.HTMLElement) {
		l := e.ChildAttrs("div", "value")
		if len(l) > 0 {
			fmt.Println()
		}
		fmt.Println(l)
	})
	fmt.Println("scrape starting...\n")
	def_collector.Visit("https://dictionary.com/browse/a")
	//
	// word_collector.OnHTML("main ul", func(e *colly.HTMLElement) {
	// 	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
	// 		word := el.ChildText("a:first-child")

	// 		var data_line []string
	// 		var definitions []string

	// 		data_line = append(data_line, word)
	// 		writer.Write(data_line)
	// 	})
	// 	fmt.Println("\nscrape done :)\n")
	// })
	// word_collector.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })
	// scraping := true
	// word_collector.OnError(func(r *colly.Response, err error) {
	// 	scraping = false
	// })
	// word_collector.Visit("https://dictionary.com/list/0")
	// for i := 0; i < len(alph); i++ {
	// 	ind := 1
	// 	for scraping {
	// 		word_collector.Visit("https://dictionary.com/list/" + alph[i] + "/" + strconv.Itoa(ind))
	// 		ind += 1
	// 	}
	// 	scraping = true
	// }
	fmt.Println("program ending...")
}
