package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello 世界")

	var foo int = 10
	var bar int = 5

	sum, product := SumAndProduct(foo, bar)
	fmt.Printf("X: %d\tY: %d\nSum: %d\nProduct: %d", foo, bar, sum, product)
}

func SumAndProduct(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return
}
