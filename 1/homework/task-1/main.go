package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(MinElement(a, b, c, d))

}

//MinElement return minimal element of 4 arguments
func MinElement(e int, f int, g int, h int) int {
	arr := [4]int{e, f, g, h}
	min := arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	return min
}
