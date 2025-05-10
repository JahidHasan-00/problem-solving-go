package main

import "fmt"

func main() {
	var (
		t int
		n int
	)
	fmt.Scan(&t)
	if t >= 1 && t <= 99 {
		for i := 0; i < t; i++ {
			fmt.Scan(&n)
			if n >= 2 && n <= 100 {
				fmt.Println(n - 1)
			}
		}
	}
}
