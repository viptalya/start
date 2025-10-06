package main

import "fmt"

//Псевдонимы
//Оператор type позволяет определять для типа псевдоним, при этом это не новый тип, а новое название для существующего

// Для типа int определяется псидвоним mile, который основывается на типе int. По сути - это новый тип
type mile int
type km int

func Distance(d mile) {
	fmt.Println(d)
}

func Demo9() {
	var d mile = 5
	Distance(d)

	//var k km = 6
	//Distance(k) ошибка, т.к. k - это тип km, а наш метод принимает тип mile
}
