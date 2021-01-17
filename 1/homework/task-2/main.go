package main

import "fmt"

func main() {
	var n int
	fmt.Println("Plase enter number to return from Fibonacci sequence:")
	fmt.Scan(&n)
	fmt.Println(fibN(n))
}

//Return n number from Fibonacci sequence.
func fibN(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
