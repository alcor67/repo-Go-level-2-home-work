// race1_with_mutex project main.go
package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

//sync.RWMutex
type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = struct{}{}

}
func (s *Set) Has(i int) (struct{}, bool) {
	s.RLock()
	defer s.RUnlock()
	st, ok := s.mm[i]
	return st, ok
}

const count = 5

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	var set = NewSet()
	var (
		counter int
		lock    sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			set.Add(counter)
			counter += 1
			fmt.Println("Add counter: ", counter)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			st, ok := set.Has(counter - 1)
			counter += 1
			fmt.Println("Has counter: ", counter, st, ok)
		}()

	}
	wg.Wait()
	fmt.Println(counter)
}
