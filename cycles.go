package main

import "fmt"

// Циклы
func Cycles() {

	//Объявления
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}

	//Типо while
	i := 1

	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	//Работа с наборами данных (range, len())
	str := "Hello"

	for index, value := range str {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}

	//Если не нужно какое-то значение, то можно написать:
	for _, value := range str {
		fmt.Printf("value: %c\n", value)
	}

	//Для перебора массива можно использовать len
	var users = [3]int{1, 2, 3}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	//break и continue понятно

	//метка для выхода из внешнего цикла

OuterLoop:
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)

			if i == 2 && j == 2 {
				fmt.Println("Выход из внешнего цикла")
				break OuterLoop
			}
		}
	}

	fmt.Println("Цикл завершен")
}
