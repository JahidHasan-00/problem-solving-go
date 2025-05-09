package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	totalCandies := num * (num + 1) / 2
	fmt.Println(totalCandies)
}
