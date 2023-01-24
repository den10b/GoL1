package main

import (
	"fmt"
	"os"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
var wg sync.WaitGroup
var mx sync.Mutex
var ch chan int
var sum int

func squareMX(num int) {
	mx.Lock()
	_, _ = fmt.Fprint(os.Stdout, sum, "+=", num, "^2\n")
	sum = sum + (num * num)
	mx.Unlock()
	wg.Done()
}
func squareChan(num int) {
	old := <-ch
	_, _ = fmt.Fprint(os.Stdout, old, "+=", num, "^2\n")
	ch <- old + (num * num)
	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum = 0
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через Mutex:\n")
	for _, num := range arr {

		go squareMX(num) //Через WG+Mutex
	}
	wg.Wait()
	_, _ = fmt.Fprint(os.Stdout, "Ответ: ", sum, "\n\n")

	sum = 0
	ch = make(chan int, len(arr))
	ch <- 0
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через chan:\n")
	for _, num := range arr {
		go squareChan(num) //Через WG+chan
	}
	wg.Wait()
	_, _ = fmt.Fprint(os.Stdout, "Ответ: ", <-ch, "\n")
	close(ch)
}
