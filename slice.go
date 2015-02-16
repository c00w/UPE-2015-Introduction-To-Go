package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := make([]byte, 10)
	fmt.Printf("Capacity of a is %d Length is %d\n", cap(a), len(a))
	b := make([]byte, 0, 9001)
	fmt.Printf("Capacity of b is %d Length is %d\n", cap(b), len(b))

	c := append(a, 'h', 'i', '!')
	c = append(c, c...)

	fmt.Printf("Capacity of c is %d Length is %d\n", cap(c), len(c))
	fmt.Printf("c is %q\n", string(c))

	resp, _ := http.Get("http://rcos.rpi.edu")
	defer resp.Body.Close()

	n, _ := resp.Body.Read(c[0:32])
	fmt.Printf("%d bytes read = %q\n", n, string(c[0:n]))
}
