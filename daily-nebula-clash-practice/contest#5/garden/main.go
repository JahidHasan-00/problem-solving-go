package main

import "fmt"

func areaOfSquare(a, b int) int {
	totalAreaOfRoad := a + b - 1
	return a*b - totalAreaOfRoad
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	totalSquareYards := areaOfSquare(a, b)
	fmt.Println(totalSquareYards)
}
