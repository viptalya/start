package main

import "fmt"

// Условия
func Terms() {

	//пример if...else
	a := 10
	b := 8

	if a > b {
		fmt.Println("a больше b")
	} else if a < b {
		fmt.Println("a меньше b")
	} else {
		fmt.Println("a равно b")
	}

	//в условии if можно объявить переменные, но они видны только внутри блока
	if x := 10; x > a {
		fmt.Println("x больше a")
	}

	//fmt.Println(x) ошибка

	//switch...case
	switch a {
	case 9:
		fmt.Println("a == 9")
	case 8:
		fmt.Println("a == 8")
	case 7:
		fmt.Println("a == 7")
	default:
		fmt.Println("значение a ?")
	}

	//fallthrough - заставляет выполнение переходить от одного случая к другому
	switch a {
	case 9:
		fmt.Println("a == 9")
	case 10:
		fmt.Println("a == 10")
		fallthrough
	case 7:
		fmt.Println("a == 7")
	default:
		fmt.Println("значение a ?")
	}

	/*
		вывод:
		a == 10
		a == 7
	*/
}
