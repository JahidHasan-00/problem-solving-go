package main

import "fmt"

func main() {
	var (
		n int
		t int
		a int
	)

	fmt.Scan(&n, &t, &a)
	if n >= 1 && n <= 99 && n%2 != 0 {
		remain := n - (t + a)
		if t >= a+remain || a >= t+remain {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else {
		fmt.Println("Invalid input")
	}
}
