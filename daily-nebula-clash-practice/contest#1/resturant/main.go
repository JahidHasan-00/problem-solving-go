package main

import "fmt"

func main() {
	var (
		n int
	)
	fmt.Scan(&n)

	if n >= 1 && n <= 100 {
		x := 800 * n
		y := 200 * (n / 15)
		fmt.Println(x - y)
	}

}
