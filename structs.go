package main

import (
	"fmt"
	"github.com/CurryBoy/RPI-Computing-2013-Introduction-To-Go/geometry"
	"math/rand"
	"time"
)

func main() {
	rgen := rand.New(rand.NewSource(time.Now().Unix()))
	uniform := func(a, b float32) float32 {
		return rgen.Float32()*(b-a) - a
	}
	p1 := geometry.Point{uniform(-100, 100), uniform(-100, 100)}
	p2 := p1.Add(geometry.Point{uniform(-100, 100), uniform(-100, 100)})

	fmt.Printf("Point 1: %v\n", p1)
	fmt.Printf("Point 2: %v\n", p2)

	line := geometry.NewLine(p1, p2)
	fmt.Printf("Point 1 and Point 2 are on the line:\n%v", line)
}
