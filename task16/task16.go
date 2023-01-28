package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка

func quickSort(myArray []int, indxs ...int) { //Можно передавать массив не через ссылку, тк меняем его внутренности
	var begin, end int
	if len(indxs) == 0 {
		begin = 0
		end = len(myArray) - 1
	} else {
		begin = indxs[0]
		end = indxs[1]
	}
	if begin < end {
		pivot := partition(myArray, begin, end)
		quickSort(myArray, begin, pivot-1)
		quickSort(myArray, pivot+1, end)
	}

}
func partition(myArray []int, begin, end int) int {
	mid := (begin + end) / 2
	myArray[end], myArray[mid] = myArray[mid], myArray[end]
	pivot := myArray[end]
	i := begin
	for j := begin; j < end; j++ {
		if myArray[j] < pivot {
			myArray[i], myArray[j] = myArray[j], myArray[i]
			i++
		}
	}
	myArray[i], myArray[end] = myArray[end], myArray[i]
	return i
}

func main() {
	a := []int{1, 2, 3, 5, 10, 4, 8}
	fmt.Println(a)
	quickSort(a)
	fmt.Println(a)
}
