package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b int
	fmt.Scan(&a)
	b = int(math.Pow(a, 2))
	fmt.Println(b)
}
