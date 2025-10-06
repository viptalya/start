package main

import "fmt"

//Анонимные функции
//Анонимная функция - используется для определения действия там, где оно применяется, если в коде больше нигде не будет необходимости

// Функция демонстрации анонимной функции
func Demo3() {

	//Анонимная функция
	f := func(x, y int) int {
		return x + y
	}

	fmt.Println(f(5, 6))
}

// Анонимная функция как аргумент функции
func Demo4() {
	AnonimAction(5, 6, func(x, y int) int { return x + y })
	AnonimAction(1, 10, func(x, y int) int { return x - y })
}

// Функция получает переменные и функцию, которая будет работать с переменными
func AnonimAction(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}

//Анонимная функция как результат функции

func Demo5() {

	f := SelectAnonimFn(1)
	fmt.Println(f(5, 6))
}

// Анонимные функции являются результатом выполнения этой функции
func SelectAnonimFn(n int) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }
	} else if n == 2 {
		return func(x, y int) int { return x * y }
	} else {
		return func(x, y int) int { return x - y }
	}
}
