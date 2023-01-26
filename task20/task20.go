package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»
func reverseStrings(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseStrings(input[1:]), input[0])
}
func main() {
	myString := "snow dog sun"
	fmt.Println(myString)

	wordsArr := strings.Fields(myString)
	newString := strings.Join(reverseStrings(wordsArr), " ")
	fmt.Println(newString)

	newString = ""
	for i := len(wordsArr) - 1; i >= 0; i-- {
		newString = newString + wordsArr[i] + " "
	}
	fmt.Println(newString)
}
