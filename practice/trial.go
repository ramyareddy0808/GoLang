package main

import (
	"fmt"
)

// type person struct {
// 	firstName string
// 	lastName  string
// 	Age       int
// }

func main() {

	// s := "appleappleapple"
	// count := 0
	// for i := 0; i < len(s); i++ {
	// 	if string(s[i]) == "a" {
	// 		count++
	// 	}
	// 	fmt.Println(count)
	// }

	// s := "appleappleapple"
	// countOfa := strings.Count(s, "a")
	// fmt.Println(countOfa)

	// myName := person{firstName: "Ramya", lastName: "Reddy", Age: 20}
	// fmt.Println(myName)

	colors := map[string]string{
		"red":    "apple",
		"yellow": "banana",
		"orange": "orange",
	}

	fmt.Println(colors["yellow"])

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number%2 != 0 {
			fmt.Println(number, " is not divisible by 2")
		} else {
			fmt.Println(number, "is divisible by 2")
		}

	}
}
