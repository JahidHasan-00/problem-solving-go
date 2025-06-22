package main

import "fmt"

func main() {
	var t, x, d int
	fmt.Scan(&t)
	d = 5
	if t >= 1 && t <= 10 {
		for i := 0; i < t; i++ {
			fmt.Scan(&x)
			if x >= 1 && x <= 10 {
				fmt.Println(2 * x * d)
			}
		}
	}
}
