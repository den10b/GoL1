package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество.

func main() {
	arr1 := []string{"cat", "cat", "dog", "cat", "tree"}
	mmap := make(map[string]struct{})
	for _, j := range arr1 {
		mmap[j] = struct{}{} //Создаем map, где ключами будут являться значения из массива
	}
	var result []string
	for key := range mmap {
		result = append(result, key)
	}
	fmt.Println(result)
}
