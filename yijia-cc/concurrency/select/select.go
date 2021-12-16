package main

import (
	"concurrency/log"
	"time"
)

func Query(ids []string, logger *log.Logger) string {
	// Buffer added to reslove deadLock issue when receiving from resultCh
	resultCh := make(chan string, 1)
	for _, id := range ids {
		go func(id string) {
			select {
			case resultCh <- <-fetch(id, logger):
				logger.Log(log.Info, "Done fetching %s", id)
			default:
				logger.Log(log.Info, "Default: %s", id)
			}
		}(id)
	}

	time.Sleep(time.Duration(5) * time.Second)
	logger.Log(log.Info, "Start receiving the first query result")
	return <-resultCh
}

func fetch(id string, logger *log.Logger) chan string {
	var delay int
	ch := make(chan string)

	switch id {
	case "Apple":
		delay = 2
	case "Banana":
		delay = 1
	default:
		delay = 3
	}

	go func(id string) {
		logger.Log(log.Debug, "Start fetching %s", id)
		time.Sleep(time.Duration(delay) * time.Second)
		ch <- id
		close(ch)
		logger.Log(log.Debug, "End fetching %s", id)
	}(id)

	return ch
}

func main() {
	logger := log.NewLogger(log.Debug)
	ids := []string{"Apple", "Banana", "Watermelon", "Blueberry"}
	logger.Log(log.Info, "Final result: %v\n", Query(ids, &logger))
}
