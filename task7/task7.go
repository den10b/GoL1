package main

import (
	"fmt"
	"log"
	"sync"
)

// Реализовать конкурентную запись данных в map.

var wg sync.WaitGroup
var ch chan mapThread

type mapThread struct {
	sync.RWMutex
	numbers map[int]int
}

func (mapp *mapThread) mapThreadWrite(i int) {
	defer wg.Done()
	mapp.Lock()
	mapp.numbers[i] = i //тк приисполнении данного кода объект изменяется, то мы должны вызывать только из 1 горутины единовременно горутин, поэтому здесь используется Lock
	mapp.Unlock()
	log.Printf("%d added\n", i)
}
func (mapp *mapThread) mapThreadRead(i int) {
	defer wg.Done()
	mapp.RLock()
	fmt.Println(mapp.numbers[i]) //тк приисполнении данного кода объект не изменяется, то мы можем  вызывать её параллельно из любого количества горутин, поэтому здесь используется RLock вместо Lock
	mapp.RUnlock()
}
func (mapp *mapThread) mapThreadReadAll() {
	defer wg.Done()
	mapp.RLock()
	log.Println(mapp.numbers) //тк приисполнении данного кода объект не изменяется, то мы можем  вызывать её параллельно из любого количества горутин, поэтому здесь используется RLock вместо Lock
	mapp.RUnlock()
}

func mapThreadWriteChan(i int) {
	mapp2 := <-ch //тк канал не буферизированный каждая горутина ждет пока другая горутина положит map в канал
	mapp2.numbers[i] = i
	wg.Done()
	ch <- mapp2 // таком случае единовременно запись в канал может исполнять только одна горутина
}

func main() {

	mapp := mapThread{
		numbers: map[int]int{},
	}
	n := 100
	wg.Add(n + 2)
	for i := 0; i < n; i++ {
		go mapp.mapThreadWrite(i)
		if (i == 50) || (i == 95) {
			go mapp.mapThreadReadAll() //конкурентная запись с чтением через RWMutex
		}
	}
	wg.Wait()
	log.Print(mapp)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go mapp.mapThreadRead(i) //конкурентное чтение
	}
	wg.Wait()

	mapp = mapThread{
		numbers: map[int]int{},
	}
	wg.Add(n)
	ch = make(chan mapThread)
	for i := 0; i < n; i++ {
		go mapThreadWriteChan(i) //конкурентная запись с через chan
	}
	ch <- mapThread{
		numbers: map[int]int{},
	}
	wg.Wait()
	mapp = <-ch
	fmt.Println(mapp)

}
