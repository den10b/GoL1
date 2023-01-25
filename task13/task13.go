package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 14
	b := 1
	fmt.Println("a=", a, " b=", b)
	a = a + b
	b = a - b //Стандартный способ обмена через вычисления
	a = a - b
	fmt.Println("a=", a, " b=", b)
	b, a = a, b //Приравнивание одновременно а к b, b к а
	fmt.Println("a=", a, " b=", b)
}
