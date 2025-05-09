/****
	#Problem No:
	-> 1929. Concatenation of Array

	#Problem Type: Easy
	#Follow or open to the link: https://leetcode.com/problems/concatenation-of-array/description/

****/

package main

import "fmt"

func getConcatenation(nums []int) []int {
	l := len(nums)
	ans := make([]int, l*2)

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+l] = nums[i]
	}
	return ans
}

func main() {
	nums := []int{1, 3, 2, 1}
	res := getConcatenation(nums)
	fmt.Println(res)
}

/**Another way to solve this problem:

	func getConcatenation(nums []int) []int {
    ans := append(nums, nums...);
    return ans;
}
**/
