package main

import (
	"fmt"
	"os"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

var ch1 chan int //Канал для записи чисел из массива
var ch2 chan int //Канал для записи результата операции
var wg sync.WaitGroup

func producer(arr []int) {
	defer wg.Done()
	for _, j := range arr {
		ch1 <- j
	}
	close(ch1) //При закрытии канала из него все еще можно получить оставшиеся значения
}
func executer() {
	defer wg.Done()
	for i := range ch1 { //При получении всех значений из канала и его закрытии - выходим из цикла
		ch2 <- i * 2
	}
	close(ch2)
}
func reciever() {
	defer wg.Done()
	for i := range ch2 {
		_, _ = fmt.Fprint(os.Stdout, i, "\n")
	}
}
func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch1 = make(chan int, len(arr))
	ch2 = make(chan int, len(arr))
	wg.Add(3)
	go producer(arr)
	go executer()
	go reciever()
	wg.Wait()
}
