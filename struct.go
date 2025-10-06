package main

import "fmt"

//Структуры
//Структура - составной тип данных для представление объектов. Для определения структуры применяются ключевые слова type и struct
/*

type имя_структуры struct {
	поле1 тип_поля1
	поле2 тип_поля2
	...
	полеN тип_поляN
}

*/

// Создание структуры
type person struct {
	name string
	age  int
}

// Работа с обычной структурой
func Demo10() {

	var tom person
	fmt.Println(tom) //{ 0} т.к. мы не присвоили полям никаких значение, то они получают значения по умолчанию

	//Инициализация структуры
	var tom1 person = person{"Tom1", 21}
	fmt.Println(tom1) //{Tom1 21}

	//Обращение к полям структуры
	tom1.age = 23
	fmt.Println(tom1.name) //Tom1
	fmt.Println(tom1.age)  //23

}

// Объявление анонимной структуры
var tom struct {
	name string
	age  int
}

// Работа с анонимной структурой
func Demo11() {
	tom.name = "Tom Anon"
	tom.age = 22
	fmt.Println(tom) //{Tom Anon 22}
}

// Анонимные поля структуры
type personAnon struct {
	string
	int
}

//Если у нас более одного поля одинакового типа, то использовать как анонимное можно только одно, для остальных необходимо указать имя поля
/*

type person struct {
	string
	// string Ошибка: string redeclared - other declaration of string
	company string
	int
}

*/

// Работа с анонимными полями структуры
func Demo12() {
	tom := personAnon{"TomAnon", 25}
	fmt.Println(tom) //{TomAnon 25}
}

//Указатели на структуру
//Как и с переменными можно создать указатель, который хранит адрес структуры

type person1 struct {
	string
	int
}

// Работа с указателями на структуру
func Demo13() {
	tom := person1{"Tom", 24}

	var pTom *person1 = &tom
	//указатель на структуру
	fmt.Println(*pTom)

	//обращение к полям через указатель
	fmt.Println(pTom.string)

	//обращение к полям структуры через разыменовывание указателя
	(*pTom).int = 34
	fmt.Println((*pTom).string)
	fmt.Println((*pTom).int)
}

//Функция new()
//Выделяет пространство памяти в соответсвии с полями структуры и возвращает указатель на нее. Изначально все поля имют значение по умолчанию

type person2 struct {
	name string
	age  int
}

func Demo14() {
	tom := new(person2)
	fmt.Println(*tom) //{ 0}

	tom.name = "Tom"
	tom.age = 21

	fmt.Println(*tom) //{Tom 21}

	//анонимная структура с new
	bob := new(struct {
		name string
		age  int
	})

	bob.name = "Bob"
	bob.age = 45
	fmt.Println(*bob)
}

// Копирование структуры
func Demo15() {
	tom := person2{"Tom", 42}

	tomas := tom
	tomas.age = 18

	fmt.Println(tom.age)
	fmt.Println(tomas.age)
}

//Передача структуры в функцию

func Increment_Age(user *person2) {
	user.age += 1
}

func Demo16() {
	tom := person2{"tom", 15}

	//в функцию передается адрес структуры
	Increment_Age(&tom)

	fmt.Println(tom) // {tom 16}
}

// Сравнение структур
func Demo17() {
	tom := person2{"Tom", 17}
	bob := person2{"Bob", 58}
	tomas := person2{"Tom", 17}

	fmt.Println(tom == bob)   //false
	fmt.Println(tom == tomas) //true
}
