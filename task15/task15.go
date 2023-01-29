package main

import (
	"fmt"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.

var justString string

func createHugeString(size int) string {
	res := ""
	for i := 0; i < size; i++ {
		res += string(byte(111 + i))
	}
	return res
}

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100] в таком случае мы получаем не строку в 100 символов, а <=100, тк
	//при обращении по индексу мы получаем фиксированное количество байт, в то время как
	//некоторые символы кодируются несколькими байтами
	justString = v[:100] //неправильный вариант
	badR := []rune(justString)
	fmt.Println("old len = ", len(badR)) //Выводит 64

	justString = string([]rune(v)[:100]) //правильный вариант
	newR := []rune(justString)
	fmt.Println("new len = ", len(newR)) //Выводит 100

}

func main() {
	someFunc()
}
