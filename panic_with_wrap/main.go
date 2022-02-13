// error project main.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Calling f()")
	f()
	fmt.Println("Returned normally from f()")
}
func f() {

	defer func() {
		panicRecover := recover()
		//приведение типа Type assertions err, ok := panicRecover.(error)
		if err, ok := panicRecover.(error); err != nil {
			if !ok {
				fmt.Println("panicRecover type mismatch with err type")
				os.Exit(1)
			}

			//https://golangify.com/date-time-layout-2006-01-02

			timeOfError := time.Now().Format(time.RFC1123)
			err = errors.Wrapf(err, "time of error: %v ", timeOfError)
			fmt.Println("Recovered in f(): ", err)
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
