package main

import "fmt"

func main() {
	var a, b, c, d, o, p int

	fmt.Scan(&a, &b, &c, &d)
	o, p = MinMaxElement(a, b, c, d)
	fmt.Printf("Min element is: %v\nMax element is: %v\n", o, p)

}

//MinMaxElement return minimal and maximal element of 4 arguments
func MinMaxElement(e int, f int, g int, h int) (int, int) {
	arr := [4]int{e, f, g, h}
	min := arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	max := arr[0]
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return min, max
}
