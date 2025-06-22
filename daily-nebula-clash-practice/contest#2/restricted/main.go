/***
	#Restricted:
	-->Problem Statement:
		You are given two integers A and B as the input. Output the value of A+B.
		However, if A+B is 10 or greater, output error instead.

***/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a >= 1 && b <= 9 && a+b < 10 {
		fmt.Println(a + b)
	} else {
		fmt.Println("error")
	}
}
