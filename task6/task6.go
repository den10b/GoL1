package main

import (
	"context"
	"fmt"
	"github.com/matryer/runner"
	"sync"
	"time"
)

// 6. Реализовать все возможные способы остановки выполнения горутины.

var wg sync.WaitGroup

func main() {
	ch := make(chan string, 6)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Хватит работать")
		for {
			v, ok := <-ch //При получении данных из канала получаем также его состояние
			if !ok {      //Если канал закрыт - горутина завершает работу
				return
			}
			fmt.Println(v)
		}
	}()
	ch <- "Работа"
	ch <- "Опять работа"
	close(ch)
	wg.Wait()
	fmt.Println("Finish")

	ch = make(chan string, 6)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Хватит работать")
		for {
			for v := range ch { //как только канал закрывается - выходим из цикла и останавливаем горутину
				fmt.Println(v)
			}
			return
		}
	}()
	ch <- "Работа"
	ch <- "Опять работа"
	close(ch)
	wg.Wait()
	fmt.Println("Finish")

	ch = make(chan string, 6)
	done := make(chan struct{}) //используем переменную done как семафор для остановки горутины
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("Хватит работать")
		for {
			select {
			default:
				fmt.Println("Работаем...")
			case <-done: //проверяем наличие чего-нибудбь в канале done
				close(ch)
				return
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
	done <- struct{}{} //заносим в канал
	wg.Wait()
	fmt.Println("Finish")

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		defer wg.Done()
		defer fmt.Println("Хватит работать")
		for {
			select {
			case <-ctx.Done(): // при получении каоманды происходит завершение гроутины
				return
			default:
				fmt.Println("Работаем...")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	cancel() //Используем функцию отмены, полученную при создании контекста

	wg.Wait()
	fmt.Println("Finish")

	wg.Add(1)
	task := runner.Go( //Используем библиотеку runner, в которой предусмотрена остановка горутин извне
		func(shouldStop runner.S) error {
			defer wg.Done()
			defer fmt.Println("Хватит работать")
			for {
				fmt.Println("Работаем...")
				if shouldStop() {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			return nil
		})
	time.Sleep(3 * time.Second)
	task.Stop()
	wg.Wait()
	fmt.Println("Finish")
}
