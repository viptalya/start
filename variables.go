package main

import "fmt"

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
