package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в
//конкурентной среде. По завершению программа должна выводить итоговое
//значение счетчика.

type Counter struct {
	count      int
	sync.Mutex //так как при записи читать не обязательно можем использовать обычный мьютекс
}
type CounterChan struct {
	count int
}

var wg sync.WaitGroup
var counter Counter
var ch chan CounterChan

func increment() {
	counter.Lock()
	counter.count = counter.count + 1
	counter.Unlock()
	wg.Done()
}

func incrementChan() {
	count := <-ch
	count.count = count.count + 1
	wg.Done()
	ch <- count
}

func main() {
	counter = Counter{count: 0}
	N := 5
	wg.Add(N)
	for i := 0; i < N; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println(counter.count)

	ch = make(chan CounterChan, 1)
	counterChan := CounterChan{count: 0}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go incrementChan()
	}
	ch <- counterChan
	wg.Wait()
	counterChan = <-ch
	fmt.Println(counterChan.count)

}
