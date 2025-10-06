package main

import "fmt"

//Объявление функции
/*
func имя_функции (список_параметров) типы_возвращаемых_значений {
	тело_функции
}
*/

func FuncSumWithParams(x, y int) {
	fmt.Println(x + y)
}

// Функция с неопределенным количеством параметров
func FuncWithNParams(numbers ...int) {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("sum =", sum)
}

//Передача среза в функцию
/*
Чтобы передать срез в функцию необходимо после аргумента-массива указать многоточие

FuncWithNParams([]int{1, 2, 3}...)

или

var nums = []int{5, 6, 7}
FuncWithNParams(nums...)
*/

//Возвращение результата из функции
/*
func имя_функции (список_параметров) тип возвращаемого_значения {
	тело_функции
	return возвращаемое_значение
}
*/

func FuncSumWithReturn(x, y int) int {
	return x + y
}

// Именованные возвращаемые результаты
func FuncSumWithReturnZ(x, y int) (z int) {
	z = x + y
	return
}

// Возвращение нескольких результатов
func FuncReturnNParams(x, y int, firstName, lastName string) (z int, fullName string) {
	z = x + y
	fullName = firstName + " " + lastName
	return
}

func Calculate() {
	a := FuncSumWithReturn(5, 6)
	fmt.Println(a)

	b := FuncSumWithReturnZ(11, 10)
	fmt.Println(b)

	age, name := FuncReturnNParams(11, 14, "Vitalii", "Podoliak")
	fmt.Printf("Возраст = %v\nПолное имя = %v\n", age, name)
}

//Типы функции
/*
Тип функции складывается из списка параметров и списка типов возвращаемых результатов

Например, функция

func add(x, y int) int {
	return x + y
}

имеет тип - func(int, int) int

А функция

func display(s string) {
	fmt.Printls(s)
}

имеет тип - func(string)

Используется для того, чтобы мы могли переменной присвоить в качестве типа функцию
*/

// Функция для демонтсрации
func Demo() {

	//Присваиваем переменной а тип функции FuncSum - func(int, int) int
	a := FuncSum
	fmt.Println(a(3, 5))

	//Теперь переменная а указывает на функцию FuncMultiply, у которой тип тоже func(int, int) int
	a = FuncMultiply
	fmt.Println(a(3, 5))

	//a = FuncDisplay - ошибка, т.к. невозможно переменной с типом func(int, int) int присвоить значение с типом func(string)

	//Присваиваем переменной s тип func(string)
	s := FuncDisplay
	s("Текст")
}

func FuncSum(x, y int) int {
	return x + y
}

func FuncMultiply(a, b int) int {
	return a * b
}

func FuncDisplay(s string) {
	fmt.Println(s)
}

//Функции как параметры других функций
//Функция может передаваться в качестве параметра в другую функцию

func Demo1() {
	slice := []int{-2, 4, 3, -1, 7, -4, -23}

	//получим сумму чисел по условию функции isEven
	sumOfEvens := ResultSum(slice, isEven)
	fmt.Println(sumOfEvens)

	//получим сумму чисел по условию функции isPositive
	sumOfPositive := ResultSum(slice, isPositive)
	fmt.Println(sumOfPositive)
}

// Функция проверки четного числа
func isEven(n int) bool {
	return n%2 == 0
}

// Функция проверки положительного числа
func isPositive(n int) bool {
	return n > 0
}

// Функция, которая принимает срез и в соответствии с заданным условием выводит результат
func ResultSum(nums []int, criteria func(int) bool) int {
	result := 0

	for _, num := range nums {
		if criteria(num) {
			result += num
		}
	}

	return result
}

//Функция как результат другой функции
//Функция также может возвращаться из другой функции в качестве результата

func Demo2() {
	f := selectFn(1)
	fmt.Println(f(3, 4))

	f = selectFn(3)
	fmt.Println(f(3, 4))
}

// Функция суммирования
func sum1(x, y int) int {
	return x + y
}

// Функция вычитания
func subtract1(x, y int) int {
	return x - y
}

// Функция умножения
func multiply1(x, y int) int {
	return x * y
}

// Функция, которая возвращает функцию типа func(int, int) int, в зависимости от n
func selectFn(n int) func(int, int) int {

	switch n {
	case 1:
		return sum1
	case 2:
		return subtract1
	default:
		return multiply1
	}
}

//Если определить функцию с большой буквы, то она будет доступна для других пакетов, если определить с маленькой, то другие пакеты ее не увидят.
