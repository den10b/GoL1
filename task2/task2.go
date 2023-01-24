package main

import (
	"fmt"
	"os"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

var wg sync.WaitGroup
var mx sync.Mutex
var ch chan int

func squarePrintWG(num int) {
	square := num * num
	_, _ = fmt.Fprint(os.Stdout, square, "\n")
	wg.Done()
}
func squarePrintMX(num int) {
	mx.Lock()
	square := num * num
	_, _ = fmt.Fprint(os.Stdout, square, "\n")
	mx.Unlock()
	wg.Done()
}
func squarePrintChan(num int) {
	ch <- num * num
	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через WaitGroup:\n")
	for _, num := range arr {
		go squarePrintWG(num) //Через WaitGroup
	}
	wg.Wait()

	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через Mutex:\n")
	for _, num := range arr {
		go squarePrintMX(num) //Через WG+Mutex
	}
	wg.Wait()

	ch = make(chan int, len(arr))
	wg.Add(len(arr))
	_, _ = fmt.Fprint(os.Stdout, "Через chan:\n")
	for _, num := range arr {
		go squarePrintChan(num) //Через WG+chan
	}
	for _, _ = range arr {
		_, _ = fmt.Fprint(os.Stdout, <-ch, "\n")
	}
	close(ch)
	wg.Wait()
}
