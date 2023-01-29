package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.

func check1(myString string) bool { //
	mapp := make(map[int32]struct{})
	for _, st := range strings.ToLower(myString) { //По очереди проверяем каждый символ
		if _, ok := mapp[st]; ok { //оператор if может начинаться с инструкции, которая будет выполнена перед проверкой условия
			return false //если ключ найден, значит такой символ уже был в строке
		} else {
			mapp[st] = struct{}{} //если ключ не найден, значит такого символа еще не было, проверяем дальше
		}

	}
	return true
}
func main() {
	stringA := "abcdутыв"
	stringB := "abCdefAaf"
	stringC := "aabcd"
	fmt.Println(check1(stringA))
	fmt.Println(check1(stringB))
	fmt.Println(check1(stringC))
}
