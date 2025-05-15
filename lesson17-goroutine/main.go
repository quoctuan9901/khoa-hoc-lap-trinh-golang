package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func heavyTask(wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for i := 0; i < 100e8; i++ {
		sum += i
	}
}

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println("So CPU:", numCPU)

	runtime.GOMAXPROCS(numCPU)

	start := time.Now()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go heavyTask(&wg)
	}

	wg.Wait()

	fmt.Println("Tong thoi gian: ", time.Since(start))
}
