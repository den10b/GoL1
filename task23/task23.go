package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.

func del(a *[]int, i int) {
	if i >= len((*a)) {
		panic("индекс больше длины массива")
	}
	*a = append((*a)[:i], (*a)[i+1:]...) //Разделяем изначальный массив на 2 слайса, затем обьединяем их
}

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice)
	del(&mySlice, 4)
	fmt.Println(mySlice)

}
