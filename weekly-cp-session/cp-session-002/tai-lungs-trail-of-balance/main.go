package main

import "fmt"

func main() {
	var t, n, a int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		var s []int
		sum := 0
		for j := 0; j < n; j++ {
			fmt.Scan(&a)
			s = append(s, a)
			sum += a
		}

		if sum%len(s) == 0 {
			fmt.Println(len(s))
		} else {
			fmt.Println(sum)
		}
	}
}
