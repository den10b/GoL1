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
	for _, st := range strings.ToLower(myString) {
		if _, ok := mapp[st]; ok {
			return false
		} else {
			mapp[st] = struct{}{}
		}

	}
	return true
}
func main() {
	stringA := "abcd"
	stringB := "abCdefAaf"
	stringC := "aabcd"
	fmt.Println(check1(stringA))
	fmt.Println(check1(stringB))
	fmt.Println(check1(stringC))
}
