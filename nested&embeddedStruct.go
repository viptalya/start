package main

import "fmt"

//Вложенные и встроенные структуры
/*

Вложенные структуры (nested struct) определяются как именованные поля внутри другой структуры, и доступ к их полям осуществляется с использованием поля внешней структуры

Встроенные структуры (embedded struct) определяются без именя поля, что дает прямой доступ к своим полям из внешней структуры

*/

//Вложенные структуры

type personNested struct {
	name string
	age  int
}

type accountNested struct {
	login    string
	password string

	personInfo personNested
}

func Demo18() {
	tom := accountNested{
		login:    "tom@mail.com",
		password: "123",
		personInfo: personNested{
			name: "Tom",
			age:  18,
		},
	}

	fmt.Println(tom) // {tom@mail.com 123 {Tom 18}}

	//Обращение к полям вложенной структуры
	fmt.Println(tom.personInfo.name) //Tom
}

//Встроенная структура

type personEmbedded struct {
	name string
	age  int
}

type accountEmbedded struct {
	login    string
	password string
	personEmbedded
}

func Demo19() {
	bob := accountEmbedded{
		"bob@mail.com",
		"567",
		personEmbedded{"Bob", 56},
	}

	fmt.Println(bob) //{bob@mail.com 567 {Bob 56}}

	//Обращение к полям встроенной структуры
	fmt.Println(bob.personEmbedded.name) //Bob
}
