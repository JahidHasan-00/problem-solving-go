package main

import "fmt"

func search(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}

	return -1

}

func main() {
	//Slice literal:
	nums := []int{-1, 0, 2, 4, 6, 8}
	target := 3

	//Passed arguments in function and received a return value:
	result := search(nums, target)
	fmt.Println(result)

}
