package main

import (
	"fmt"
	"time"
)

func HardWorker(state int, input <-chan []int, output chan<- int, quit chan<- int) {
	work := <-input
	for _, v := range(work) {
		time.Sleep(50 * time.Millisecond)
		output <- state + v
	}
	quit <- 1
}

func main() {
	work := []int{1,2,3,4,5}
	outCh := make(chan int)
	quitCh := make(chan int)
	for i := 0; i < 10; i++ {
		inCh := make(chan []int)
		go HardWorker(10*i, inCh, outCh, quitCh)
		inCh <- work
	}
	counter := 0
	for counter < 10 {
		select {
		case v := <-outCh:
			fmt.Println(v)
		case <- quitCh:
			counter++
		}
	}
}
