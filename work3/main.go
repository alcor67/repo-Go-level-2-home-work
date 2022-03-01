//Simulate a “race” situation, and check the program
//for the presence of a “race”
package main

import (
	"fmt"
	"sync"
)

func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	mm := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{7, 8, 9, 10},
	}
	nn := [][]int{
		{10, 11, 12},
		{13, 14, 15},
		{16, 17, 18},
		{19, 20, 21},
	}

	ll := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	llCh := make(chan [3][3]int)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go multipleMatrix(mm, nn, llCh)
	ll = <-llCh
	close(llCh)
	go func([3][3]int) {
		defer wg.Done()
		//time.Sleep(1 * time.Second)
		for _, l := range ll {
			fmt.Println(l)
		}
	}(ll)
	wg.Wait()
}

func multipleMatrix(mm, nn [][]int, llCh chan [3][3]int) {
	var wg = sync.WaitGroup{}
	var ll [3][3]int
	for i := 0; i < len(mm); i++ {
		for j := 0; j < len(mm); j++ {
			for k := 0; k < len(mm[1]); k++ {
				wg.Add(1)
				go func(i, j, k int) {
					defer wg.Done()
					ll[i][j] = ll[i][j] + mm[i][k]*nn[k][j] //строка mm[i] столбец mm[k] + строка nn[k] столбец nn[j]
				}(i, j, k)
			}
		}
	}
	wg.Wait()
	llCh <- ll
}
