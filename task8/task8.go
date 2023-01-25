package main

import (
	"fmt"
	"math"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.
func main() {
	num := int64(62174938754656182)
	fmt.Printf("Дано число %d\n", num)
	i := 27
	val := 1 //Желаемое значение i-того бита

	fmt.Printf("Бит на позиции %d должен стать %d\n", i, val)
	var result int64

	switch val {
	case 0:
		temp := int64(math.Pow(2, 64)) - 1
		mask := int64(math.Pow(2, float64(i)))
		mask = temp - mask //Создаем маску состоящую из 1 и одного 0 на месте бита, который надо сделать 0
		fmt.Printf("%65b\n", mask)
		fmt.Printf("%65b\n", num)
		result = mask & num //Применяем бинарное И, которое изменяет только целевой бит
	case 1:
		mask := int64(math.Pow(2, float64(i))) //Создаем маску состоящую из 0 и одной 1 на месте бита, который надо сделать 1
		fmt.Printf("%65b\n", mask)
		fmt.Printf("%65b\n", num)
		result = mask | num //Применяем бинарное ИЛИ, которое изменяет только целевой бит

	}

	fmt.Printf("%65b\n", result)
	fmt.Printf("Резулльтат %d\n", result)
}
