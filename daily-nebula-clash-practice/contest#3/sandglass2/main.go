/***
	##D - Sandglass :

	-->Problem Statement:
		We have a sandglass that runs for X seconds. The sand drops from the upper bulb at a rate of 1 gram per second. 		That is, the upper bulb initially contains X grams of sand.

		How many grams of sand will the upper bulb contains after t seconds?
***/

package main

import (
	"fmt"
)

func main() {
	var x, t int
	fmt.Scan(&x, &t)
	if x >= 1 && x <= 1000000000 && t >= 1 && t <= 1000000000 {
		if x > t {
			fmt.Println(x - t)
		} else {
			fmt.Println(0)
		}
	}
}
