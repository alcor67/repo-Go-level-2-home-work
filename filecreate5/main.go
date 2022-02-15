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

		retErr = func() (retErrIn error) {
			fileName := strconv.Itoa(i) + ".txt"

			//todo error provocation by func os.Create()
			//fileName = string(i) + ".txt"

			f, err := os.Create(fileName)

			if err != nil {
				err = errors.Wrapf(err, "error in func CreateFile() os.Create()")
				return err // sets retErrIn = err
			}

			defer closeFile(f, fileName)

			fmt.Printf("file %s Created\n", fileName)

			//todo error provocation by func f.Close()
			//defer closeFile(f, fileName)

			message := "message Nr: " + strconv.Itoa(i)

			_, err = fmt.Fprintln(f, message)
			if err != nil {
				err = errors.Wrapf(err, "error in func CreateFile() fmt.Fprintln()")
				return err // sets retErrIn = err
			}

			fmt.Printf("file %s Printed in file\n", fileName)

			return nil
		}()
	}
	return //retErr
}

func closeFile(f *os.File, fileName string) {
	if err := f.Close(); err != nil {
		err = errors.Wrapf(err, "error in func closeFile() f.Close()")
		fmt.Println(err)
	} else {
		fmt.Printf("file %v Closed\n", fileName)
	}
}
