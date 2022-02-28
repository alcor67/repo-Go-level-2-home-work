// Matrix_matrix_mult_new_goroutine_new project main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		mm := [][]int{
			{11, 12, 13},
			{21, 22, 23},
			{31, 32, 33},
		}
		nn := [][]int{
			{11, 12, 13},
			{21, 22, 23},
			{31, 32, 33},
		}
	*/

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

	ll = multipleMatrix(mm, nn)
	for _, l := range ll {
		fmt.Println(l)
	}
}

func multipleMatrix(mm, nn [][]int) [3][3]int {
	var wg = sync.WaitGroup{}
	var ll [3][3]int

	for i := 0; i < len(mm); i++ {
		for j := 0; j < len(mm); j++ {
			for k := 0; k < len(mm[1]); k++ {
				wg.Add(1)
				go func(i, j, k int) {
					defer wg.Done()
					ll[i][j] = ll[i][j] + mm[i][k]*nn[k][j] //строка mm[i] столбец mm[k] + строка nn[k] столбец nn[j]
					//fmt.Printf("goroutine nr. i: %v, j: %v, k: %v, ll[%v][%v]=%v\n", i, j, k, i, j, ll[i][j])
				}(i, j, k)
			}
		}
	}
	wg.Wait()
	return ll
}

/*
[160 170 180]
[334 356 378]
[508 542 576]
*/
