package main

import "fmt"

// Массивы
func Arrays() {

	//var arr [размерность]тип_данных
	//var arr1 [3]int = [3]int{1, 2, 3}
	//var arr2 := [...]int{1} длина массива вычисляется из количества переданных элементов

	//многомерные
	//arr3 := [размерность1][размерность2]..тип_данных{значения}
	//arr4 := [3][2] int {{1, 2}, {4, 5}, {7, 8}}

	//длина массива
	arr5 := [2]int{0, 1}

	fmt.Printf("Значения массива: %v \n", arr5)
	fmt.Printf("Длина массива: %v", len(arr5))

	arr6 := [3][2]int{{1, 2}, {2, 4}}

	fmt.Printf("\nЗначения массива: %v \n", arr6)
	fmt.Printf("Длина массива: %v\n", len(arr6))

	//копирование массива
	arr7 := [2]int{1, 2}
	arr8 := arr7

	arr8[0] = 10

	fmt.Println(arr7)
	fmt.Println(arr8)
}
