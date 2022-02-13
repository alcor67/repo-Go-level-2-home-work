// error project main.go
package main

import (
	"fmt"
	"os"
	"time"

	"runtime/debug"

	"github.com/pkg/errors"
)

//объявляем структуру содержащую текст ошибки и время ее возникновения
type ErrorWithTime struct {
	message string
	errTaim string
	trace   string
}

//инициируем структуру, т.е. заносим данные в соответствующие поля структуры
func myNewErrorMethod(message string) error {
	return &ErrorWithTime{
		message: message,
		//https://golangify.com/date-time-layout-2006-01-02
		errTaim: string(time.Now().Format(time.RFC1123)),
		trace:   string(debug.Stack()),
	}
}

//добавляем метод Error(), возвращающий строку, где ресивером будет служить
//указатель на объявленную выше структуру ErrorWithTime
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("%s\ntime: %s\ntrace:\n%s", e.message, e.errTaim, e.trace)
}

func main() {
	fmt.Println("Calling f()")
	f()
	fmt.Println("Returned normally from f()")
}
func f() {

	defer func() {
		panicRecover := recover()
		//приведение типа Type assertions panicRecover, ok := panicRecover.(error)
		if panicRecover, ok := panicRecover.(error); panicRecover != nil {
			if !ok {
				fmt.Println("panicRecover type mismatch with runtime.errorString type")
				os.Exit(1)
			}
			//fmt.Printf("Type assertions panicRecover: %T\n", panicRecover)
			/*
				преобразование ошибки err в строку panicRecover.Error()
				причем без приведения типа panicRecover, ok := panicRecover.(error)
				метод panicRecover.Error() не работает
				panicRecover.Error undefined (type interface {} is interface with no methods)
				и оборачивание с фразой found error:
			*/

			panicRecover = errors.Wrap(myNewErrorMethod(panicRecover.Error()), "found error")

			fmt.Println("Recovered in f(): ", panicRecover)
		}
	}()

	fmt.Println("Calling g()")
	g(0)
	fmt.Println("Returned normally from g()")
}
func g(i int) {

	j := 0
	for ; i < 10; i++ {
		if i > 3 {
			fmt.Println("Panicking in g() !")
			_ = i / j
		}
		defer fmt.Println("Defer in g()", i)
		fmt.Println("Printing in g()", i)
	}
}
