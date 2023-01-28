package main

import (
	"fmt"
)

//Реализовать пересечение двух неупорядоченных множеств

func main() {
	arr1 := []int{5, 1, 2, 7, 9, 0, 4}
	arr2 := []int{10, 5, 19, 2, 11, 1, 89}
	mmap := make(map[int]struct{})
	var result []int
	for _, j := range arr1 {
		mmap[j] = struct{}{} //Создаем map, где ключами будут являться значения из первого массива
	}
	for _, j := range arr2 {
		if _, ok := mmap[j]; ok { //Там, где значения ключам не присвоены будет false
			result = append(result, j) //Там, где значения ключам присвоены будет true, соответственно число содержится в первом массиве
		}
	}
	fmt.Println(result)
	mmap = make(map[int]struct{})
	result = []int{}
	for _, j := range arr1 { //Более медленный вариант. Решение в лоб
		for _, k := range arr2 {
			if j == k {
				mmap[j] = struct{}{}
			}
		}
	}
	for key := range mmap {
		result = append(result, key)
	}
	fmt.Println(result)
}
