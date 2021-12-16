package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := time.After(2 * time.Second)
	timeout := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		timeout <- struct{}{}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("finished processing")
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
