// error project main.go
package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/pkg/errors"
)

func main() {
	err := CreateFile(5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func CreateFile(i int) (retErr error) {

	for i := 0; i < 5; i++ {

		fileName := strconv.Itoa(i) + ".txt"
		//error provocation by func os.Create()
		//fileName = string(i) + ".txt"

		f, err := os.Create(fileName)
		if err != nil {
			err = errors.Wrapf(err, "error in func CreateFile() os.Create()")
			return err // sets retErr = err
		}
		fmt.Printf("file %s Created\n", fileName)

		//error provocation by func f.Close()()
		//defer closeFile(f, fileName)
		message := "message Nr: " + strconv.Itoa(i)
		_, err = fmt.Fprintln(f, message)
		if err != nil {
			err = errors.Wrapf(err, "error in func CreateFile() fmt.Fprintln()")
			return err // sets retErr = err
		}
		fmt.Printf("file %s Printed in file\n", fileName)

		// вообще-то это то же самое, что прсто closeFile(f, fileName)
		//т.е. вызов closeFile идет в этом же самом месте после выхода из анонимной функции
		//тогда именно в такой ситуации в чем смысл defer?
		//а вот ошибку из-под defer никак не вытащить.
		func() { defer closeFile(f, fileName) }()
	}
	return nil
}

func closeFile(f *os.File, fileName string) {
	if err := f.Close(); err != nil {
		err = errors.Wrapf(err, "error in func closeFile() f.Close()")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("file %v Closed\n", fileName)
	}
}
