// goroutines project main.go
package main

import (
	"fmt"

	"github.com/alcor67/repo-Go-level-2-home-work/Homework-4/workerpool"
)

func main() {

	finalNumber := workerpool.WorkerPool(1000)
	fmt.Println(finalNumber)
}
