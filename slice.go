package main

import (
	"fmt"
	"foo"
	"net/http"
)

func main() {
	a := make([]byte, 10)
	foo.Fill(a, 10)
	fmt.Printf("Capacity of a is %d\n", cap(a))

	c := append(a, 'h', 'i', '!')
	c = append(c, c...)

	fmt.Printf("Capacity of c is %d\n", cap(c))
	fmt.Printf("C is \"%s\"\n", string(c))

	resp, _ := http.Get("http://rcos.rpi.edu")
	defer resp.Body.Close()

	n, _ := resp.Body.Read(c[0:20])
	fmt.Printf("32 bytes read = \"%s\"\n", string(c[0:n]))
}
