package main

import (
	"fmt"
	"start/lasagna"
	_ "start/messages/messages_en"
	_ "start/messages/messages_ru"
)

func main() {
	fmt.Println(lasagna.ElapsedTime(5, 5))
}
