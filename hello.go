package main

import "fmt"

func main() {
	Demo17()
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

//Анонимные функции
//Анонимная функция - используется для определения действия там, где оно применяется, если в коде больше нигде не будет необходимости

//Функция демонстрации анонимной функции
func Demo3() {

	//Анонимная функция
	f := func(x, y int) int {
		return x + y
	}

	fmt.Println(f(5, 6))
}

//Анонимная функция как аргумент функции
func Demo4() {
	AnonimAction(5, 6, func(x, y int) int { return x + y })
	AnonimAction(1, 10, func(x, y int) int { return x - y })
}

//Функция получает переменные и функцию, которая будет работать с переменными
func AnonimAction(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}

//Анонимная функция как результат функции

func Demo5() {

	f := SelectAnonimFn(1)
	fmt.Println(f(5, 6))
}

//Анонимные функции являются результатом выполнения этой функции
func SelectAnonimFn(n int) func(int, int) int {
	if n == 1 {
		return func(x, y int) int { return x + y }
	} else if n == 2 {
		return func(x, y int) int { return x * y }
	} else {
		return func(x, y int) int { return x - y }
	}
}

/*
Замыкание (closure) - объект функции, который запоминает свое лексическое окружение, даже когда не выполняется в своей области видимости.

Замыкание включает в себя:

1. Внешняя функция, которая определяет некоторую область видимости и в которой опеределены некоторые параметры и переменные - лексическое окружение
2. Переменные и параметры (лексическое окружение), которые определены во внешней функции
3. Вложенная функция, которая использует переменные и параметры внешней функции

Замыкания реализуются с помощью анонимных функций.
*/
func Demo6() {

	//fn = inner, т.к. функция Outer возвращает функцию inner
	fn1 := Outer()

	fn1() //6
	fn1() //7

	fn2 := Multiply2(5)
	result1 := fn2(5)
	fmt.Println(result1)

	result2 := fn2(6)
	fmt.Println(result2)
}

//Внешняя функция Outer
func Outer() func() {

	//Некоторая переменная - лексическое окружение функции inner
	n := 5

	//Вложенная функция
	inner := func() {

		//Действия с переменной n
		n += 1
		fmt.Println(n)
	}

	return inner
}

//Еще пример замыкания
func Multiply2(n int) func(int) int {
	return func(m int) int { return n * m }
}

/*
Описание функции Multiply2:

Вызов функции приводит в вызову внутренней функции func(m int) int { return n * m }, которая запоминает окружение, в нашем случае переменную n
При вызове функции Multiply2 определяется переменная func, которая и представляет собой замыкание, т.е. объединяет функцию и окружение, в котором была создана.

Т.е. result1 - это замыкание, которое содержит и внутреннюю функцию func(m int) int { return n * m } и параметр n, который существует во время создания замыкания.

Важно не запутаться в параметрах. При определении замыкания:

fn2 := Multiply2(5) - число 5 передается для параметра n функции Multiply2
result1 := fn2(5) - число 5 передается для параметра m внутренней функции func(m int) int { return n * m }
*/

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

//Пустой указатель
//Если указателю не присвоен адрес какого-либо объекта, то его значение будет nil и лучше добавить проверку на nil
func NilIndicator() {
	var p *int

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("нет адреса объекта")
	}
}

//Функция new
//С помощью функции new(type) можно создать безымянный объект указанного типа
func NewIndicator() {

	p := new(int)
	fmt.Println(*p) //0 - значение по умолчанию
	*p = 8
	fmt.Println(*p) //8
}

//Указатели на указатели
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

//Массивы указателей
//Можем создать массив указателей, которые содержат адрес значений одного и того же типа данных
func ArrIndicator() {
	var a, b, c int = 1, 2, 3

	p_nums := [4]*int{&a, &b, nil, &c}

	for _, value := range p_nums {
		if value != nil {
			fmt.Println(*value)
		}
	}
}

//Указатели как параметры функции
func ChangeValue(x *int) {
	*x = (*x) * (*x)
}

func Demo7() {
	d := 5
	fmt.Println(d) //5
	ChangeValue(&d)
	fmt.Println(d) //25
}

//Указатель как результат функции
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

//Псевдонимы
//Оператор type позволяет определять для типа псевдоним, при этом это не новый тип, а новое название для существующего

//Для типа int определяется псидвоним mile, который основывается на типе int. По сути - это новый тип
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

//Создание структуры
type person struct {
	name string
	age  int
}

//Работа с обычной структурой
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

//Объявление анонимной структуры
var tom struct {
	name string
	age  int
}

//Работа с анонимной структурой
func Demo11() {
	tom.name = "Tom Anon"
	tom.age = 22
	fmt.Println(tom) //{Tom Anon 22}
}

//Анонимные поля структуры
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

//Работа с анонимными полями структуры
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

//Работа с указателями на структуру
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

//Копирование структуры
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

//Сравнение структур
func Demo17() {
	tom := person2{"Tom", 17}
	bob := person2{"Bob", 58}
	tomas := person2{"Tom", 17}

	fmt.Println(tom == bob)   //false
	fmt.Println(tom == tomas) //true
}
