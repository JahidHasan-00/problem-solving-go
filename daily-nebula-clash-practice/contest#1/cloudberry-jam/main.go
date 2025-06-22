package main

import "fmt"

func main() {
	var t, n int
	fmt.Scan(&t)
	if t > 0 && t <= 10000 {
		for i := 0; i < t; i++ {
			fmt.Scan(&n)
			if n >= 1 && n <= 100000000 {
				fmt.Println(n * 2)
			}
		}
	}
}
