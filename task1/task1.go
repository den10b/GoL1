package main

import (
	"fmt"
	"math"
)

// Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

type Human struct {
	name   string
	age    int
	height int
	weight float64
}

func (h Human) BMI() float64 {
	return h.weight / math.Pow(float64(h.height)/100, 2)
}

type Action struct {
	Action string
	Human  //Все методы и поля типа Human автоматически делаются доступными через тип Action
}
type Action2 struct {
	Action   string
	Executor Human //Все методы и поля типа Human будут доступными через Action.Executor
}

func (a Action2) BMI() float64 {
	return a.Executor.BMI() //Менее удобный аналог встраивания методов
}

func main() {
	dima := Human{"Дима", 20, 185, 65}
	action1 := Action{"Sing", dima}
	fmt.Printf("Доступ к полю родительской структуры: %s\nВызов метода без обращения к родительской структуре %f\n", action1.name, action1.BMI())

	den := Human{"Денис", 21, 180, 60}
	action2 := Action2{"Dance", den}
	fmt.Printf("Доступ к полю родительской структуры: %s\nВызов метода с обращением к родительской структуре %f\n", action2.Executor.name, action2.Executor.BMI())
}
