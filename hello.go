package main

import (
	"start/messages/messages_en"
	"start/messages/messages_ru"
)

func main() {
	messages_en.Hello_en()
	messages_en.Bye_en()

	messages_ru.Hello_ru()
	messages_ru.Bye_ru()
}
