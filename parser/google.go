package parser

import (
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
)

const (
	GoogleSearchURL = "https://www.google.com.ua/?hl=ru"
	GoogleSearchXPath = `//*[@id="tsf"]`

)

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func GoogleSearchParser() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)

		doc, err := htmlquery.LoadURL("https://www.bing.com/search?q=golang")
		if err != nil {
			panic(err)
		}
		// Find all news item.
		for i, n := range htmlquery.Find(doc, "//ol/li") {
			a := htmlquery.FindOne(n, "//a")
			fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
		}
	}

}
