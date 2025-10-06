package main

import "fmt"

//Методы
/*

Метод - функция связанная с определенным типом. В определении метода необходимо указать получателя - параметр того типа, к которому прикрепляется метод

func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращающих_результатов) {
	тело_метода
}
*/

// создаем срез из строк
type library []string

// создаем метод Pring, получателем которого является срез строк library
func (l library) Print() {
	for _, val := range l {
		fmt.Println(val)
	}
}

func Demo20() {
	lib := library{"book1", "book2", "book3"}

	//поскольку Print представляет метод для типа library, то вызвать его можем только у объекта типа library
	lib.Print()
}

//Методы структур

// объявляем структуру
type personMethod struct {
	name string
	age  int
}

// создаем метод и в качестве получателя указываем структуру personMethod
func (p personMethod) Print() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func Demo21() {
	tom := personMethod{"Tom", 22}
	tom.Print()
}

// Методы указателей
// Метод принимает указатель на структуру personMethod
func (p *personMethod) update_age(new_age int) {
	p.age = new_age
}

func Demo22() {
	tom := personMethod{"Tom", 22}
	fmt.Println(tom.age)

	tom.update_age(44)
	fmt.Println(tom.age)
}
