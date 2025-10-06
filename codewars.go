package main

import "fmt"

// Задача CountingSheep
func Codewars() {
	arrSheeps := [...]bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}

	countSheeps := 0

	for i := 0; i < len(arrSheeps); i++ {
		if arrSheeps[i] == true {
			countSheeps++
		}
	}

	fmt.Print(countSheeps)
}
