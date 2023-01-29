package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в
//конкурентной среде. По завершению программа должна выводить итоговое
//значение счетчика.

type CounterMX struct {
	count      int
	sync.Mutex //так как при записи читать не обязательно можем использовать обычный мьютекс
}
type CounterChan struct {
	count int
}

var wg sync.WaitGroup
var counter CounterMX
var ch chan CounterChan

func incrementMX() {
	counter.Lock() // Пока одна горутина считает - остальные ждут
	counter.count = counter.count + 1
	counter.Unlock()
	wg.Done()
}

func incrementChan() {
	count := <-ch
	count.count = count.count + 1
	wg.Done() //Так как канал небуферизированный отправка в канал после Done во избежании дедлока
	ch <- count
}

func main() {
	counter = CounterMX{count: 0}
	N := 5
	wg.Add(N)
	for i := 0; i < N; i++ {
		go incrementMX()
	}
	wg.Wait()
	fmt.Println(counter.count)

	ch = make(chan CounterChan) //Канал, в который кладем один экземпляр счетчика, таким образом лишь одна горутина имеет к нему единовременный доступ
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
