package main

import "fmt"

func gcd(b, a int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func main() {
	var (
		t int
		a int
		b int
	)
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&b, &a)
		res := gcd(a, b)
		fmt.Println(res)
	}
}
