package main

import "fmt"

//Указатели
//Указатель - хранит адрес переменной

func Indicator() {
	x := 4         //объявляем переменную
	var p *int     //объявляем указатель
	p = &x         //передаем указателю адрес переменной x
	fmt.Println(p) //0xc00008c060

	//Чтобы по указателю вывести результат переменной надо перед ним поставить *
	fmt.Println(*p) //4

	//Используя указатель можем изменить значение переменной по адресу
	*p = 25
	fmt.Println(x) //25
}

// Пустой указатель
// Если указателю не присвоен адрес какого-либо объекта, то его значение будет nil и лучше добавить проверку на nil
func NilIndicator() {
	var p *int

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("нет адреса объекта")
	}
}

// Функция new
// С помощью функции new(type) можно создать безымянный объект указанного типа
func NewIndicator() {

	p := new(int)
	fmt.Println(*p) //0 - значение по умолчанию
	*p = 8
	fmt.Println(*p) //8
}

// Указатели на указатели
func IndOnInd() {

	a := 4
	pa := &a
	fmt.Println(*pa) //4
	fmt.Println(pa)  //0xc00008c0a0

	pp := &pa
	fmt.Println(pp)   //0xc00008e038
	fmt.Println(*pp)  //0xc00008c0a0 содержит адрес указателя pa
	fmt.Println(**pp) //4 чтобы получить значние переменной, на которую указывает pa, то надо перед указателем pp поставить **
}

// Массивы указателей
// Можем создать массив указателей, которые содержат адрес значений одного и того же типа данных
func ArrIndicator() {
	var a, b, c int = 1, 2, 3

	p_nums := [4]*int{&a, &b, nil, &c}

	for _, value := range p_nums {
		if value != nil {
			fmt.Println(*value)
		}
	}
}

// Указатели как параметры функции
func ChangeValue(x *int) {
	*x = (*x) * (*x)
}

func Demo7() {
	d := 5
	fmt.Println(d) //5
	ChangeValue(&d)
	fmt.Println(d) //25
}

// Указатель как результат функции
func CreatePointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func Demo8() {
	p1 := CreatePointer(10)
	fmt.Println(*p1)
	p2 := CreatePointer(15)
	fmt.Println(*p2)
}
