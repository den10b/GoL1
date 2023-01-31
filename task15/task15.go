package main

import (
	"fmt"
	"strings"
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
	//некоторые символы в UTF8 кодируются несколькими байтами

	//Также т.к justString - глобальная переменная-string(слайс байт), в ней будет хранится ссылка на большой базовый массив,
	// из-за этого сборщик мусора не будет очищать его до самого завершения программы
	justString = v[:100] //неправильный вариант
	badR := []rune(justString)
	fmt.Println("old len = ", len(badR)) //Выводит 64

	temp := string([]rune(v)[:100])
	justString = strings.Clone(temp) //правильный вариант
	newR := []rune(justString)
	fmt.Println("new len = ", len(newR)) //Выводит 100

}

func main() {
	someFunc()
}
