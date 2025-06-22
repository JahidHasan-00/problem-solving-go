package main

import "fmt"

func main() {
	var t, d int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		val := true
		var c int = 0
		for i := 0; i < 7; i++ {
			fmt.Scan(&d)
			if d > 1 || d < 0 {
				val = false
			}
			c += d
		}

		if c >= 5 && val == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
