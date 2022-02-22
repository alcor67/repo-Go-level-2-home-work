// context project main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	err := Ping()
	fmt.Println("main err:", err)
}
func Ping() error {
	//Создание нового контекста на основе базового context.Background()
	ctx, cancelFunc := context.WithCancel(context.Background())
	//Перед штатным выходом из функции освобождаем ресурсы, выделенные под контекст
	defer cancelFunc()
	//создание канала по системным сигналам
	sigCh := make(chan os.Signal, 1)
	/*
		https://pkg.go.dev/os/signal
		В связи с тем, что у меня установлена ОС Windows 8.1
		у меня возникли сложности с получением сигнала SIGTERM.
		В качестве альтернативы в коде обрабатывается сигнал SIGINT или же
		как вариант os.Interrupt
	*/
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		//блокировка  по системному сигналу os.Interrupt завершающей контекст горутины
		sig := <-sigCh
		log.Println("received signal:", sig)

		//завершение контекста по cancelFunc()
		cancelFunc()
	}()
	//Создаём канал, в который долго работающая функция в горутине может сообщить результаты выполнения: ошибка или нет
	doneCh := make(chan error)
	//В отдельной горутине запускает потенциально долго работающую функцию, передав ей контекст на вход.
	go func(inCtx context.Context) {
		err := Pong(inCtx)
		doneCh <- err //ошибка err отправляется в канал doneCh
	}(ctx)

	var err error
	//С помощью select дожидаемся первого события: либо результата работы функции,
	//либо завершения контекста по cancelFunc().
	select {
	//Метод контекста Done() <-chan struct{} возвращает канал, значение в который придет при завершении всех функций, опирающийся на этот контекст.
	case <-ctx.Done():
		/*Если контекст вынужденно завершился, метод Err() error вернёт ошибку
		  context.Canceled (контекст был принудительно завершен функцией cancelFunc())
		  или context.DeadlineExceeded (контекст вынужденно завершился из-за таймаута)
		*/
		err = ctx.Err()
	// контекст штатно завершился получаем ошибку из канала doneCh
	case err = <-doneCh:
	}
	return err
}

func Pong(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		log.Println("Loop tick")
		time.Sleep(1 * time.Second)
	}
	return nil
}
