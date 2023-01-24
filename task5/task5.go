package main

import (
	"log"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться.

var ch chan int
var wg sync.WaitGroup

func Reader() {
	defer wg.Done()
	for {
		num, op := <-ch //Так как мы пользуемся каналом, при его закрытии мы можем закончить работу всех воркеров
		if op {
			log.Println(num)
		} else {
			log.Println("Reader остановлен")
			return
		}
	}
}

func main() {
	N := 10
	bTime := time.Now()
	defer log.Println("main остановлен")

	chCount := 0
	ch = make(chan int, chCount)
	go Reader()
	wg.Add(1) // Запускаем получателя

	i := 0
	delta := time.Now().Sub(bTime)
	for delta.Seconds() <= float64(N) { //Проверяем разницу между началом программы и текущим временем
		if i >= 100 {
			i = 1
		} else { //Генерируем данные для отправки в канал
			i++
		}
		ch <- i
		delta = time.Now().Sub(bTime) //Ищем разницу между началом программы и текущим временем
	}
	close(ch)
	wg.Wait()

}
