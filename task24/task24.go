package main

import (
	p "GoL1/task24/point"
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
func findDist(a *p.Point, b *p.Point) float64 {
	return math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))
}

func main() {
	// строчка вида new:=p.Point{1,2} выдаст ошибку, тк поля структуры не экспортируются

	a := p.NewPoint(1, 1)
	b := p.NewPoint(2, 2)

	fmt.Println(findDist(a, b))
}
