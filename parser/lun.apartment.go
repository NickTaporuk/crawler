package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
)

const (
	NowostroikiUri       = "https://novostroyki.lun.ua/все-новостройки-киева"
	NewBuildingCardXPath = `//*[@class="card-grid-cell"]		`
)

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func LunParser() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)

		doc, err := htmlquery.LoadURL(NowostroikiUri)
		if err != nil {
			panic(err)
		}
		// Find all news item.
		for i, n := range htmlquery.Find(doc, NewBuildingCardXPath) {
			a := htmlquery.FindOne(n, `//*[@class="card-image"]`)
			fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "src"))
			//fmt.Println(i, n)
		}
	}

}
