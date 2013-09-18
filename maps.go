package main

import (
    "fmt"
)

func main() {
    heights := make(map[string]int)
    heights["bharath"] = 5
    heights["colin"] = 6

    fmt.Println(heights)
    fmt.Println(heights["joe"])

    for k,v := range(heights) {
        fmt.Printf("%s is %d feet high\n", k, v)
    }
}
