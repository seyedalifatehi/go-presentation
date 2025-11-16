package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// for example we want to fetch users from db and it costs 100ms for each user
func FetchUser() {
	<-time.After(100 * time.Millisecond)
}

func Sum(from, to int) int {
	s := 0

	for i := from; i <= to; i++ {
		s += i
	}

	return s
}

func main() {
	// the number of cpu cores to use.
	// still 100ms -> FetchUser() function.
	// for Sum() function this causes more time and should be
	// greater than 1 to reduce time of calculating
	runtime.GOMAXPROCS(1)

	startTime := time.Now()

	// 500 ms wait and it is not good. instead we can use "goroutines"
	// FetchUser()
	// FetchUser()
	// FetchUser()
	// FetchUser()
	// FetchUser()

	// we declare WaitGroup to ensure that all of the routines execute completely. (for FetchUser() function).
	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		// just waiting and it is concurrency means that we use just 1 CPU core
		// FetchUser()

		// we use CPU and process so if the process become parallel it takes less time
		Sum(1, 20_000_000)
		wg.Done()
	}()

	go func() {
		// FetchUser()
		Sum(1, 20_000_000)
		wg.Done()
	}()

	go func() {
		// FetchUser()
		Sum(1, 20_000_000)
		wg.Done()
	}()

	go func() {
		// FetchUser()
		Sum(1, 20_000_000)
		wg.Done()
	}()

	go func() {
		// FetchUser()
		Sum(1, 20_000_000)
		wg.Done()
	}()

	wg.Wait()

	// if i comment all the above code and run this code
	// with 1 cpu core they are executed at same time
	// the execution time of the following code may be less.
	// the explanation is that we have context switching and because of that
	// we may get more time

	// Sum(1, 100_000_000)

	fmt.Println(time.Since(startTime))
}
