package main

import (
	"fmt"
	"runtime"
	"sync"
)

func WithMutex(count int) {
	wg := sync.WaitGroup{}
	var mu sync.Mutex

	result := 0

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			result++
			mu.Unlock()
		}(i)
	}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			result--
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(result)
}

func WithRace(count int) {
	wg := sync.WaitGroup{}

	result := 0

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result++
		}(i)
		if i%10 == 0 {
			runtime.Gosched()
		}
	}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			result--
		}(i)
	}
	wg.Wait()

	fmt.Println(result)
}
func WithSched(k int) {
	res := 0
	go func() {
		res += k
	}()
	for i := 0; ; i += 1 {
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}
}
func main() {
	WithRace(10000)
}
