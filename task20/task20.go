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
	return append(reverseStrings(input[1:]), input[0]) //После достижения функцией самого правого элемента в массиве - она начинает собирать строку в обратном порядке
}
func main() {
	myString := "snow dog sun"
	fmt.Println(myString)

	wordsArr := strings.Fields(myString)                     //разделяем по пробелам на массив слов
	newString := strings.Join(reverseStrings(wordsArr), " ") //обьединяем массив строк в одну с разделителем-проблом
	fmt.Println(newString)

	newString = ""
	for i := len(wordsArr) - 1; i >= 0; i-- {
		newString = newString + wordsArr[i] + " "
	}
	fmt.Println(newString)
}
