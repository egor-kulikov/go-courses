package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	if a >= 1 {
		for i := 1; i <= a; i++ {
			sum += i
		}
		fmt.Println(sum)
	} else {
		for i := 1; i >= a; i-- {
			sum += i
		}
		fmt.Println(sum)
	}
}
