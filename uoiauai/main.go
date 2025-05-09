package main

import "fmt"

func main() {
	var char string
	fmt.Scan(&char)

	switch char {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
		break
	default:
		fmt.Println("consonant")
	}
}
