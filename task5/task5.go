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
		num, op := <-ch //Так как мы пользуемся каналом, при его закрытии мы можем закончить работу получателя
		if op {
			log.Println(num)
		} else {
			log.Println("Reader остановлен")
			return
		}
	}
}

func main() {
	N := 5
	defer log.Println("main остановлен")

	ch = make(chan int)
	go Reader()
	wg.Add(1) // Запускаем получателя
	defer wg.Wait()

	i := 0
	timeout := time.After(time.Duration(N) * time.Second) //После N секунд в канал timeout посылается текущее время

	for {
		select {
		case <-timeout: //При получении чего-то из канала timeout происходит закрытие канала и выход
			close(ch)
			return
		default:
			if i >= 100 {
				i = 1
			} else { //Генерируем данные для отправки в канал
				i++
			}
			ch <- i //Отправляем данные в канал
		}
		time.Sleep(10 * time.Millisecond)
	}

	/* Один из вариантов реализации, в котором каждую итерацию цикла проверяется разница между временем начала программы и текущим
	i := 0
	bTime := time.Now()
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



	*/
}
