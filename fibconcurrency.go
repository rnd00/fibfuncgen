package main

import (
	"runtime"
	"sync"
)

func fibConcurrency(i int) {
	numCPU := runtime.NumCPU()
	var wg sync.WaitGroup
	NumberChannel := make(chan int)

	var a func(n int) int
	a = func(n int) int {
		if n <= 2 {
			return 1
		}
		return a(n-1) + a(n-2)
	}

	b := func(nChan <-chan int, wg *sync.WaitGroup) int {
		defer wg.Done()
		var r int
		for num := range nChan {
			r = a(num)
		}
		return r
	}

	// c := func()

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go b(NumberChannel, &wg)
	}
}
