package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseInts(input []int32) []int32 {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
func main() {
	s := "hello"
	var byteString []byte

	for i := len(s) - 1; i >= 0; i-- {
		byteString = append(byteString, s[i])
	}
	newString := string(byteString)
	fmt.Println(newString)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}

	fmt.Println()
	intString := []int32{}
	for _, j := range s {
		intString = append(intString, j)
	}
	for _, j := range reverseInts(intString) {
		fmt.Printf("%c", j)
	}

}
