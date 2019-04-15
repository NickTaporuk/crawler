package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
)

const (
	NowostroikiUri       = "https://novostroyki.lun.ua/жк-итальянский-квартал-киев"
	NewBuildingCardXPath = `//*[@class="card-grid-cell"]`
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

		list := htmlquery.Find(doc, NewBuildingCardXPath)

		fmt.Println("LIST:=>", list)
	}

}
