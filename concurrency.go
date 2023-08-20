package puppy

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func Concurrency() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go cfoo()
	cbar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func cfoo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("foo: %d\n", i)
	}
	wg.Done()
}

func cbar() {
	for i := 0; i < 10; i++ {
		fmt.Printf("bar: %d\n", i)
	}
}

func Mutex() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}

func Atomic() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: \t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()

		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
