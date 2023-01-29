package main

import (
	"fmt"
	"os"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

var wg sync.WaitGroup
var mx sync.RWMutex
var ch chan int
var ch2 chan int

func squarePrintWG(num int) {
	square := num * num
	_, _ = fmt.Fprint(os.Stdout, square, "\n")
	wg.Done()
}
func squarePrintMX(num int) {
	mx.RLock()
	square := num * num //Так как данные, которые мы читаем не изменяются, можно использовать RLock
	_, _ = fmt.Fprint(os.Stdout, square, "\n")
	mx.RUnlock()
	wg.Done()
}
func squarePrintChan(num int) {
	ch <- num * num
	wg.Done()
}
func squarePrintChanBuf() {
	num := <-ch
	wg.Done()
	wg.Wait()
	ch2 <- num * num

}

func main() {
	arr := [5]int{2, 4, 6, 8, 10} //Массив, а не слайс
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через WaitGroup:\n")
	for _, num := range arr {
		go squarePrintWG(num) //Через WaitGroup
	}
	wg.Wait()

	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через Mutex:\n")
	for _, num := range arr {
		go squarePrintMX(num) //Через WG+RWMutex
	}
	wg.Wait()

	ch = make(chan int) //Используя небуф. канал
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через chan:\n")
	for _, num := range arr {
		go squarePrintChan(num) //Через WG+chan
	}
	for range arr {
		_, _ = fmt.Fprint(os.Stdout, <-ch, "\n")
	}
	close(ch)
	wg.Wait()

	ch = make(chan int, 1) //Используя буф. канал
	ch2 = make(chan int)
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через буф. chan:\n")
	for _, num := range arr {
		ch <- num
		go squarePrintChanBuf()
	}
	count := 0
	for d := range ch2 {
		_, _ = fmt.Fprint(os.Stdout, d, "\n")
		count++
		if count == len(arr) {
			break //Выходим из цикла, чтобы не было дедлока
		}
	}

}
