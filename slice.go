package main
// OMIT
import ( //OMIT
    "fmt" //OMIT
    "net/http" //OMIT
) //OMIT
//OMIT
func fill(s []byte, l byte) { // OMIT
    for i:=byte(0); i < l; i++ { //OMIT
        s[i] = byte('a') + i //OMIT
    } //OMIT
} //OMIT

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
