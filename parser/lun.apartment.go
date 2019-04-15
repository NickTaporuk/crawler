package parser

import (
	"log"
	"time"
)

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func LunParser() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(33 * time.Minute)
	}
}
