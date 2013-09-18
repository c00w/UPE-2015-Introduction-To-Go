package main
import (
    "fmt"
    "net/http"
    "foo"
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

    buf := make([]byte, 128)
    n, _ := resp.Body.Read(buf[0:32])
    fmt.Printf("32 bytes read = \"%s\"\n", string(buf[0:n]))
}
