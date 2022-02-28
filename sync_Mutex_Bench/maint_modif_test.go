// sync_Mutex_Bench_Optim_Modif project main.go
package main

import (
	"fmt"
	"sync"
	"testing"
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
func (s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

//sync.RWMutex 10% запись, 90% чтение
func BenchmarkSetAddRwm10(b *testing.B) {
	defer fmt.Println("sync.RWMutex 10% запись")
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 10; i++ {
					set.Add(i)
				}
				for i := 10; i < 100; i++ {
					set.Has(i)
				}
			}
		})
	})
}

//sync.RWMutex 90% запись, 10% чтение
func BenchmarkSetAddRwm90(b *testing.B) {
	defer fmt.Println("sync.RWMutex 90% запись")
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 10; i < 100; i++ {
					set.Add(i)
				}
				for i := 0; i < 10; i++ {
					set.Has(i)
				}
			}
		})
	})
}

//sync.RWMutex 50% запись, 50% чтение
func BenchmarkSetAddRwm50a(b *testing.B) {
	defer fmt.Println("sync.RWMutex 50% запись a")
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 50; i < 100; i++ {
					set.Add(i)
				}
				for i := 0; i < 50; i++ {
					set.Has(i)
				}
			}
		})
	})
}

//sync.RWMutex 50% запись, 50% чтение
func BenchmarkSetAddRwm50(b *testing.B) {
	defer fmt.Println("sync.RWMutex 50% запись")
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 50; i++ {
					set.Add(i)
					set.Has(i + 50)
				}
			}
		})
	})
}

//sync.Mutex
type Set1 struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet1() *Set1 {
	return &Set1{
		mm: map[int]struct{}{},
	}
}

func (s *Set1) Add1(i int) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = struct{}{}
}
func (s *Set1) Has1(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

//sync.Mutex 10% запись, 90% чтение
func BenchmarkSetAddM10(b *testing.B) {
	defer fmt.Println("sync.Mutex 10% запись 90% чтение")
	var set = NewSet1()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 10; i++ {
					set.Add1(i)
				}
				for i := 10; i < 100; i++ {
					set.Has1(i)
				}
			}
		})
	})
}

//sync.Mutex 50% запись, 50% чтение
func BenchmarkSetAddM50(b *testing.B) {
	defer fmt.Println("sync.Mutex 50% запись 50% чтение")
	var set = NewSet1()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 50; i++ {
					set.Add1(i)
				}
				for i := 50; i < 100; i++ {
					set.Has1(i)
				}
			}
		})
	})
}

//sync.Mutex 90% запись, 10% чтение
func BenchmarkSetAddM90(b *testing.B) {
	defer fmt.Println("sync.Mutex 90% запись 10% чтение")
	var set = NewSet1()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 10; i < 100; i++ {
					set.Add1(i)
				}
				for i := 0; i < 10; i++ {
					set.Has1(i)
				}
			}
		})
	})
}

//sync.Map
type Set2 struct {
	sync.Map
}

func NewSet2() *Set2 {
	return &Set2{
		Map: sync.Map{},
	}
}
func (s *Set2) Add2(i int) {
	s.Map.Store(i, struct{}{})
}
func (s *Set2) Has2(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}

//sync.Map 10% запись, 90% чтение
func BenchmarkSetAddMap10(b *testing.B) {
	defer fmt.Println("sync.Map 10% запись 90% чтение")
	var set = NewSet2()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 10; i++ {
					set.Add2(i)
				}
				for i := 10; i < 100; i++ {
					set.Has2(i)
				}
			}
		})
	})
}

//sync.Map 50% запись, 50% чтение
func BenchmarkSetAddMap50(b *testing.B) {
	defer fmt.Println("sync.Map 50% запись 50% чтение")
	var set = NewSet2()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 0; i < 50; i++ {
					set.Add2(i)
				}
				for i := 50; i < 100; i++ {
					set.Has2(i)
				}
			}
		})
	})
}

//sync.Map 90% запись, 10% чтение
func BenchmarkSetAddMap90(b *testing.B) {
	defer fmt.Println("sync.Map 50% запись 50% чтение")
	var set = NewSet2()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			//b.ResetTimer()
			for pb.Next() {
				for i := 10; i < 100; i++ {
					set.Add2(i)
				}
				for i := 0; i < 10; i++ {
					set.Has2(i)
				}
			}
		})
	})
}
