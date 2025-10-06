package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	Cycles()
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

// Внешняя функция Outer
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

// Еще пример замыкания
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

//Срезы
/*

Срезы (slice) представляют последовательность элементов одного типа переменной длины.

var users1 []string

var users2 = []string{"Tom", "Alice", "Kate"}

users3 := []string{"Tom", "Alice", "Kate"}
*/

//Оператор среза и создание среза из последовательностей
/*

Чтобы создать срез из другой последовательности применяется s[i:j]

i - начальный индекс
j - конечный индекс

Если не указан начальный индекст, то по умолчанию значение будет 0, если не указан конечный индекс, то вместо него используется длина исходной последовательности

*/

// Назначение среза из другой последовательности
func Demo23() {
	arrUsers := [6]string{"Tom", "Bob", "Kate", "Alice", "Sasha", "Mike"}
	users1 := arrUsers[2:6]
	users2 := arrUsers[:5]
	users3 := arrUsers[5:]
	users4 := arrUsers[:]

	fmt.Println(users1) //[Kate Alice Sasha Mike]
	fmt.Println(users2) //[Tom Bob Kate Alice Sasha]
	fmt.Println(users3) //[Mike]
	fmt.Println(users4) //[Tom Bob Kate Alice Sasha Mike]
}

// Создание среза из другого среза
func Demo24() {
	arrUsers := [6]string{"Tom", "Bob", "Kate", "Alice", "Sasha", "Mike"}
	users1 := arrUsers[2:6]
	users2 := users1[3:]

	fmt.Println("Slice from arr", users1)
	fmt.Println("Slice from slice", users2)
}

// Обращение к элементам среза и перебор среза
func Demo25() {
	users := []string{"Tom", "Alice", "Kate"}

	//Обращение к элементу среза производится через индекс
	users[2] = "Sasha"

	//перебор среза через range
	for _, value := range users {
		fmt.Println(value)
	}
}

/*

При инициализации среза из массива, срез ссылается на массив и при изменении элемента массива также меняется значение зреза, а при изменении элемента среза изменяет его ссылаемый массив.

*/

func Demo26() {
	arrNums := [4]int{1, 2, 3, 4}
	fmt.Println("Начальный массив", arrNums)

	slice := arrNums[2:]
	fmt.Println("Начальный срез", slice)

	//Изменяем 4й элемент массива
	arrNums[3] = 6
	fmt.Println("Срез после изменения в массиве", slice)

	//изменяем первый элемент среза
	slice[0] = 10
	fmt.Println("Массив после изменения в срезе", arrNums)
}

//Длина
/*

Длина - это кол-во элементов, на которые ссылается срез. Для получения длины принимается функция len().

*/

func Demo27() {
	arr := [3]int{1, 2, 3}
	slice := arr[:1]
	fmt.Println(len(slice)) //длина среза 1 символ
}

//Емкость
/*

Емкость среза - количество элементов в базовом массиве, начиная с первого элемента среза. Грубо говоря - емкость - это на сколько срез может быть расширен. Для получения емкости вызывается функция cap()

*/
func Demo28() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(cap(slice)) //4
}

//Создание среза с помощью функции make()

/*
make() - позволяет создать срез из  нескольких элементов, которые будут иметь значение по умолчанию.

имя_среза := make([]тип_элемента_среза, длина_среза, емкость_среза)
*/
func Demo29() {
	users := make([]int, 3)
	users[0] = 1
	users[1] = 2
	users[2] = 3

	fmt.Println(users) //[1 2 3]
}

//Двумерные срезы
/*

sice2D := [][] int {
	[]int{1, 2},
	[]int{3, 4},
	[]int{5, 6},
}

*/

//Добавление в срез
/*

Для добавления в срез применяется функция append(slice, value). Первый агрумент - срез, в который надо добавить, а второй параметр - добавляемое значение

*/

func Demo30() {
	users := []string{"Tom", "Alice", "Bob"}
	users = append(users, "Sasha")

	fmt.Println(users)
}

//Удаление элемента
/*

Чтобы удалить элемент из среза необходимо использовать функцию append с указателем среза

*/
func Demo31() {
	sliceNums := []int{1, 2, 3, 4, 5}
	//удаляем 3й элемент
	sliceNums = append(sliceNums[:2], sliceNums[2+1:]...)
	fmt.Println(sliceNums)
}

//Копирование среза

/*

Копирование среза производится с помощью функции copy()

func copy(destination, source[]T) int

Первый аргумент - срез, в который копируем, второй аргумент - срез, из которого копируем. Функция возвращает целочисленное число, которое представляет количество скопированных элементов. Возвращаемое значение всегда является длиной исходного и целевого среза

*/

// Демонстрация копирования среза
func Demo32() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{}

	copy(slice2, slice1)
	fmt.Println(slice2) //[], т.к. не указана размерность slice2

	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Println(slice3) //[1 2 3] - копирует 3 элемента, т.к. размерность slice3 - 3

	slice4 := make([]int, 6)
	copy(slice4, slice1)
	fmt.Println(slice4) //[1 2 3 4 0 0] - т.к. размер slice4 больше, чем slice1, то копируются все элементы, а оставшийся размер заполняется теми, которые были до копирования
}

//Сортировка
/*

В Go есть пакет sort, который представляет несколько методов сортировки для среза:

sort.Ints() - для типа int
sort.Float64s() - для типа float64
sort.Strings() - для типа string

Эти методы сортировки всегда выводят результат в порядке возрастания

Чтобы получить сортировку по убыванию необходимо использовать метод Reverse, только для начала необходимо обернуть в тип IntSlice/Float64Slice/StringSlice
*/
//демонстрация работы сортировки
func Demo33() {
	sliceInt := []int{1, 5, 3, 8}
	sliceFloat := []float64{1.0, 7.0, 5.4, 3.9}
	sliceString := []string{"Antony", "Lion", "Bob", "Tom"}

	//Сортировка по возрастанию
	sort.Ints(sliceInt)
	sort.Float64s(sliceFloat)
	sort.Strings(sliceString)

	fmt.Println("Sort int", sliceInt)       //Sort int [1 3 5 8]
	fmt.Println("Sort float", sliceFloat)   //Sort float [1 3.9 5.4 7]
	fmt.Println("Sort string", sliceString) //Sort string [Antony Bob Lion Tom]

	//Сортировка по убыванию
	sort.Sort(sort.Reverse(sort.IntSlice(sliceInt)))
	sort.Sort(sort.Reverse(sort.Float64Slice(sliceFloat)))
	sort.Sort(sort.Reverse(sort.StringSlice(sliceString)))

	fmt.Println("Reverse sort int", sliceInt)       //Reverse sort int [8 5 3 1]
	fmt.Println("Reverse sort float", sliceFloat)   //Reverse sort float [7 5.4 3.9 1]
	fmt.Println("Reverse sort string", sliceString) //Reverse sort string [Tom Lion Bob Antony]
}

//Поиск в срезе
/*

Для поиска в срезе в пакете sort имеются функции:

sort.SearchInts()
sort.SearchFloat64s()
sort.SearchStrings()

Эти функции выполняют поиск в отсортированном срезе в порядке возрастания. Если элемент найден, функция вернет его индекс, если нет, то индекс, по которому элемент будет помещен в срез
*/

func Demo34() {
	sliceInt := []int{11, 2, 3, 78, 55}

	fmt.Println(sort.SearchInts(sliceInt, 55))  //индекс 3 в отсортированном срезе
	fmt.Println(sort.SearchInts(sliceInt, 100)) //индекс 5 - было бы записано с индексом 5
}

//Сравнение срезов
/*

Два среза считаются равными, если они хранят данные одного типа и содержат одни и те же элементы.

Сравнить срезы можно с помощью функции DeepEqual() из пакета reflect

*/

func Demo35() {
	slice1 := []int{1, 2, 3}
	slice2 := []string{"Tom", "Bob"}
	slice3 := []int{1, 2, 3}

	fmt.Println("slice1 == slice2", reflect.DeepEqual(slice1, slice2)) //false
	fmt.Println("slice1 == slice3", reflect.DeepEqual(slice1, slice3)) //true
	fmt.Println("slice2 == slice3", reflect.DeepEqual(slice2, slice3)) //false
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
