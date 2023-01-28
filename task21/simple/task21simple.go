package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере
type requestable interface {
	Request()
}

// Клиент
type client struct{}

// Запрос клиента, который должен работать как с классом target, так и с adaptee
func (c client) Request(target requestable) {
	target.Request()
}

// Класс, к которому надо адаптировать другой класс
type target struct{}

// Метод из интерфейса
func (t target) Request() {
	fmt.Println("Base request")
}

// Адаптер
type adapter struct{ adaptee }

func (adapter adapter) Request() {
	adapter.adaptee.SpecificRequest()
}

// Адаптируемый класс
type adaptee struct{}

func (adaptee adaptee) SpecificRequest() {
	fmt.Println("Specific Request")
}

func main() {
	myClient := client{}
	myClient.Request(target{})
	ad := adapter{}
	myClient.Request(ad)
}
