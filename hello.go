package main

import (
	"fmt"
	"reflect"
)

func main() {
	Demo9()
}

//Карты (map)
/*

Карта (map) представляет собой структуру данных состоящую из пар "ключ-значение", каждый элемент имеет уникальный ключ

map[K]V, где K - тип ключа, а V - тип значения

map[тип_ключей]тип_значений {
	ключ1: значение1,
	ключ2: значение2,
	...
	ключN: значениеN
}

var people map[string]int - ключ типа string, а значение типа int
*/

func Demo36() {

	//создание пустой карты
	var people1 map[string]int
	fmt.Println(people1) //map[]

	people2 := map[string]int{
		"Tom":   22,
		"Bob":   23,
		"Sasha": 24,
	}

	fmt.Println(people2) //map[Bob:23 Sasha:24 Tom:22]

	//ключи чувствительны к регистру, поэтому могут храниться одинаковые ключи, но с разными регистрами
	people3 := map[string]int{
		"Tom": 22,
		"tom": 23,
		"Bob": 24,
		"bob": 25,
	}

	fmt.Println(people3) //map[Bob:24 Tom:22 bob:25 tom:23]
}

//Обращение к элементам карты
//Для обращение к элементам нужно применять ключи

func Demo37() {
	people := map[string]int{
		"tom":   23,
		"bob":   25,
		"alice": 21,
		"sam":   15,
	}

	fmt.Println(people["tom"])
	fmt.Println(people["alice"])

	people["alice"] = 29
	fmt.Println(people["alice"])
}

//Проверка наличия ключа
/*

Для проверки наличия элемента по определенному ключу можно применить выражение if

if i, ok := map_name[key_name]; ok {
	fmt.Println(i)
}

Если значение по заданному ключу имеется в карте, то переменная ok будет равна true, а переменная i будет хранить полученное значение.
Если переменная ok равна false, то значения по ключу в карте нет

*/

func Demo38() {

	people := map[string]int{
		"tom":   32,
		"bob":   22,
		"sam":   45,
		"alice": 3,
	}

	if val, ok := people["Sam"]; ok {
		fmt.Println(val)
	}

	//Перебор значений
	for key, val := range people {
		fmt.Println(key, val)
	}

	/*
		tom 32
		bob 22
		sam 45
		alcie 3
	*/
}

//Создание карты с помощью функции make
/*

Функция make() создает пустую хеш-таблицу

people := make(map[string]int)

*/

// Добавление и удаление элементов
func Demo39() {
	people := map[string]int{"tom": 1, "bob": 2}
	fmt.Println(people) //map[bob:2 tom:1]

	//добавление выполняется установкой значения по новому ключу
	people["sam"] = 3
	fmt.Println(people) //map[bob:2 sam:3 tom:1]

	//для удаления применяется встроенная функция delete(map, key) - первый параметр - карта, второй - ключ, по которому удалится элемент
	delete(people, "tom")
	fmt.Println(people) //map[bob:2 sam:3]
}

//Сравнение карт
/*

Также как и в срезе, для карт используется функция DeepEqual из пакета reflect. Она возвращает true, если обе карты имеют одинаковые ключи и одинаковые связанные значения ключей, иначе false

*/
func Demo40() {
	people1 := map[string]int{"tom": 1, "bob": 2}
	people2 := map[string]int{"Tom": 1, "bob": 2}
	people3 := map[string]int{"tom": 1, "bob": 2}

	fmt.Println("people1 == people2", reflect.DeepEqual(people1, people2)) //false
	fmt.Println("people2 == people3", reflect.DeepEqual(people2, people3)) //false
	fmt.Println("people1 == people3", reflect.DeepEqual(people1, people3)) //true
}
