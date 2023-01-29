package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func mySleep(second int) {
	toSleep := time.Second * time.Duration(second) //time.After() возвращает канал, в который будет послано текущее время, после того как пройдет указанный duration
	<-time.After(toSleep)                          //пока это не случится горутина ждет...
}
func mySleep2(second int) {
	bTime := time.Now()
	delta := time.Now().Sub(bTime)
	for delta.Seconds() <= float64(second) { //Проверяем разницу между началом горутины и текущим временем
		delta = time.Now().Sub(bTime) //Ищем разницу между началом программы и текущим временем
	}
}

func main() {
	fmt.Println(time.Now())
	mySleep(3)
	fmt.Println(time.Now())
	mySleep2(3)
	fmt.Println(time.Now())

}
