package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("result", queryDBs(100))
}

func queryDBs(numDB int) int {
	done := make(chan struct{})
	result := make(chan int, numDB)
	for i := 0; i < numDB; i++ {

		go func(i int) {
			select {
			case result <- query(done, i):
				// 1)
			default:
				// 1)

			}
		}(i)
	}
	return <- result
}

func query(done chan struct{}, i int) int {
	tm := time.Duration(1 + rand.Intn(5)) * time.Second
	fmt.Println(i, tm)
	time.Sleep(tm)
	return i
}