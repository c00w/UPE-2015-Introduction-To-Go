package main

import (
	"fmt"
	"time"
)

func Printer(str string, out chan<- int) {
	for i := 0; i < 10; i ++ {
		fmt.Println(str)
		time.Sleep(50 * time.Millisecond)
	}
	out <- 1
}

func main() {
	quit := make(chan int)
	go Printer("foo",quit)
	go Printer("bar",quit)
	
	<- quit
	<- quit
}
