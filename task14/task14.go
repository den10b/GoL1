package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

func SimpleDecider(a interface{}) {
	fmt.Printf("Тип: %T\n", a) //Используем Printf
}
func SwitchDecider(a interface{}) {
	switch v := a.(type) { //Используем встроенный type-switch
	case chan int:
		fmt.Printf("Тип: Integer channel\n")
	case int:
		fmt.Printf("Тип: Integer\n")
	case bool:
		fmt.Printf("Тип: Boolean\n")
	case string:
		fmt.Printf("Тип: String\n")
	default:
		fmt.Printf("Тип: %T\n", v)
	}
}
func ReflectDecider(a interface{}) {
	at := reflect.TypeOf(a).Kind() //Используем пакет Reflect
	fmt.Printf("Тип: %s\n", at)

}

func main() {
	var a chan int
	SimpleDecider(a)
	SwitchDecider(a)
	ReflectDecider(a)

}
