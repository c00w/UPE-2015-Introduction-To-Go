// Source: http://tour.golang.org/#54 OMIT
package main 
import ( 
	"fmt" 
    "os" 
) 

type Reader interface {
    Read(b []byte) (n int, err error)
}
type Writer interface {
	Write(b []byte) (n int, err error)
}
type ReadWriter interface {
    Reader
    Writer
}

func main() {
    var w Writer
    w = os.Stdout // os.Stdout implements Writer
    fmt.Fprintf(w, "hello, writer\n")
}
