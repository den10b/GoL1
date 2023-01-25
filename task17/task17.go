package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(myArray []int, key int, indxs ...int) (int, bool) { //Рекурсивная функция
	var begin, end int
	if len(indxs) == 0 {
		begin = 0
		end = len(myArray) - 1
	} else {
		begin = indxs[0]
		end = indxs[1]
	}
	if (begin >= end) && (myArray[begin] != key) {
		return 0, false
	}
	mid := (begin + end) / 2
	midElement := myArray[mid]
	if midElement == key {
		return mid, true
	} else {
		if midElement > key {
			return binarySearch(myArray, key, begin, mid-1)
		}
		if midElement < key {
			return binarySearch(myArray, key, mid+1, end)
		}
	}

	return 0, false
}
func binarySearchLoop(myArray []int, key int) (int, bool) { //Не рекурсивная функция
	begin := 0
	end := len(myArray) - 1

	for begin <= end {
		mid := (begin + end) / 2
		midElement := myArray[mid]

		if midElement < key {
			begin = mid + 1
		} else if midElement > key {
			end = mid - 1
		} else {
			return mid, true // ключ найден
		}
	}
	return 0, false // ключа нет
}

func main() {
	myArr := []int{1, 3, 5, 7, 9, 11, 15}
	toFind := 7

	notLoopKey, notLoopFind := binarySearch(myArr, toFind)
	fmt.Println(notLoopKey, notLoopFind)

	loopKey, loopFind := binarySearchLoop(myArr, toFind)
	fmt.Println(loopKey, loopFind)
}
