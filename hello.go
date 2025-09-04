package main

import "fmt"

func main() {
	Terms()
}

//Объявление переменных
func Variables() {

	const n string = "Const c"

	var a string
	a = "Variable a"

	var b string = "Variable b"

	c := "Variable c"

	d, e := "Variable d", "Variable e"

	var (
		f = "Variable f"
		g = "Variable g"
	)

	variablesArr := [...]string{a, b, c, d, e, f, g, n}

	for i := 0; i < len(variablesArr); i++ {
		fmt.Println(variablesArr[i])
	}
}

//Арифметические операции
func Operations() {
	a := 1.0
	b := 10.0

	multiplication := a * b
	summ := a + b
	subtraction := a - b
	division := a / b

	//конвертация тип(переменная)
	remains := int(a) % int(b)

	fmt.Println("Умножение", multiplication)
	fmt.Println("Сумма", summ)
	fmt.Println("Вычитание", subtraction)
	fmt.Println("Деление", division)
	fmt.Println("Деление с остатком", remains)
}

//Задача CountingSheep
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

//Массивы
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

//Условия
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
