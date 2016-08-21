package main

import (
	"fmt"
    "math3d"
)



func main() {
	var v = math3d.Vertex{23.0, 42.17}
	var v2 = v.Scale(2.0)
	fmt.Println("Yo !", v, *v2, v.LengthSq())
}
