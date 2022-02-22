// goroutines project main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	singalStop := make(chan os.Signal, 1)
	/*
		https://pkg.go.dev/os/signal
		В связи с тем, что у меня установлена ОС Windows 8.1
		у меня возникли сложности с получением сигнала SIGTERM.
		В качестве альтернативы в коде обрабатывается сигнал SIGINT или же
		как вариант так же os.Interrupt
	*/
	signal.Notify(singalStop, syscall.SIGINT)

	go func() {

		for {
			log.Println("Loop tick")
			time.Sleep(time.Second)
		}
	}()
	//блокировка завершения package main
	sig := <-singalStop
	log.Printf("received signal: %v", sig)
}
