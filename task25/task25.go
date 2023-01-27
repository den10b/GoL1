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
func main() {
	fmt.Println(time.Now())
	mySleep(3)
	fmt.Println(time.Now())
}
