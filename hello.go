package main

import "fmt"

func main() {
	Demo2()
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

//Циклы
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

func EvenNubmer() {
	arrNums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(arrNums); i++ {
		if arrNums[i]%2 == 0 {
			fmt.Printf("index: %d, value: %d - четное\n", i, arrNums[i])
		}
	}
}

//Объявление функции
/*
func имя_функции (список_параметров) типы_возвращаемых_значений {
	тело_функции
}
*/
func FuncSumWithParams(x, y int) {
	fmt.Println(x + y)
}

//Функция с неопределенным количеством параметров
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

//Именованные возвращаемые результаты
func FuncSumWithReturnZ(x, y int) (z int) {
	z = x + y
	return
}

//Возвращение нескольких результатов
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

//Функция для демонтсрации
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

//Функция проверки четного числа
func isEven(n int) bool {
	return n%2 == 0
}

//Функция проверки положительного числа
func isPositive(n int) bool {
	return n > 0
}

//Функция, которая принимает срез и в соответствии с заданным условием выводит результат
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

//Функция суммирования
func sum1(x, y int) int {
	return x + y
}

//Функция вычитания
func subtract1(x, y int) int {
	return x - y
}

//Функция умножения
func multiply1(x, y int) int {
	return x * y
}

//Функция, которая возвращает функцию типа func(int, int) int, в зависимости от n
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
