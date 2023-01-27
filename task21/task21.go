package main

//Реализовать паттерн «адаптер» на любом примере

type Client struct {
}

func (c Client) Request(target Target) {
	target.Request()
}

// класс, к которому надо адаптировать другой класс
type Target struct {
}

func (t Target) Request() {

}

// Адаптер
type Adapter struct {
	Target
	Adaptee
}

func (adapter Adapter) Request() {
	adapter.Adaptee.SpecificRequest()
}

// Адаптируемый класс
type Adaptee struct {
}

func (adaptee Adaptee) SpecificRequest() {

}

func main() {

}
