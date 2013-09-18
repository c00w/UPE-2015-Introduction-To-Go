package main

import (
    "fmt"
    "net/http"
)

func fill(s []byte, l byte) {
    for i:=byte(0); i < l; i++ {
        s[i] = byte('a') + i
    }
}

func main() {
    a := make([]byte, 10)
    fill(a, 10)
    fmt.Printf("Capacity of a is %d\n", cap(a))

    c := append(a, 'h')
    c = append(c, 'i')
    c = append(c, '!')

    c = append(c, c...)

    fmt.Printf("Capacity of c is %d\n", cap(c))
    fmt.Printf("C is \"%s\"\n", string(c))

    resp, _ := http.Get("http://rcos.rpi.edu")

    buf := make([]byte, 128)

    n, _ := resp.Body.Read(buf[0:32])
    fmt.Printf("32 bytes read = \"%s\"\n", string(buf[0:n]))
}
