package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.
var ch chan int
var wg sync.WaitGroup

func Work() {
	defer wg.Done()
	for {
		num, op := <-ch //Так как мы пользуемся каналом, при его закрытии мы можем закончить работу всех воркеров
		if op {
			log.Println(num)
		} else {
			log.Println("Worker остановлен")
			return
		}
	}
}

func main() {
	defer log.Println("main остановлен")
	ch = make(chan int)
	var wCount int
	fmt.Println("Введите кол-во Worker'ов")
	_, err := fmt.Scanf("%d\n", &wCount)
	if err != nil {
		panic(err)
	}
	chCount := 2
	c := make(chan os.Signal, chCount)
	signal.Notify(c, os.Interrupt)
	for i := 0; i < wCount; i++ {
		go Work()
		wg.Add(1) // Запускаем N рабочих
	}

	i := 0
	loop := true
	for loop == true {
		select {
		case <-c:
			close(ch)
			loop = false
		default:
			if i >= 100 {
				i = 1
			} else { //Генерируем данные для отправки в канал
				i++
			}
			ch <- i
		}
	}
	wg.Wait()

}