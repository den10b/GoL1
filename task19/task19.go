package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseInts(input []rune) []rune {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0]) //После достижения функцией самого правого символа в строке - она начинает собирать строку в обратном порядке
}
func main() {
	s := "hello"
	var runes []rune

	for i := len(s) - 1; i >= 0; i-- {
		runes = append(runes, rune(s[i]))
	}
	newString := string(runes) //Собираем rune-массив задом наперед и создаем из него строку
	fmt.Println(newString)

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}

	fmt.Println()
	intString := []rune{}
	for _, j := range s {
		intString = append(intString, j)
	}
	for _, j := range reverseInts(intString) {
		fmt.Printf("%c", j)
	}

}
