/**
	#Between Two Integers:

	Problem Statement:
	You are given three integers A, B and C. Determine whether C is not less than A and not greater than B.

**/

package main

import "fmt"

func main() {
	var (
		a, b, c int
	)
	fmt.Scan(&a, &b, &c)

	if (a >= -100 && a <= 100) && (b >= -100 && b <= 100) && (c >= -100 && c <= 100) {
		if c >= a && c <= b {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
