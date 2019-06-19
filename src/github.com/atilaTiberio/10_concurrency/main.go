package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
func main() {

	// waitingGroup()
	//raceCondition()
	atomicOperation()
}

func atomicOperation(){
	var counter int64
	const gs = 100
	wg.Add(gs)
	for i:=0; i<gs ; i++{
		go func(){
			atomic.AddInt64(&counter,1)
			runtime.Gosched()
			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Goroutines:",runtime.NumGoroutine())
	fmt.Println(counter)

}
func raceCondition(){
		counter := 0
		const gs = 100
		var mu sync.Mutex
		wg.Add(gs)
		for i:=0; i<gs ; i++{
			go func(){
				mu.Lock()
				v :=counter
				runtime.Gosched()
				v++
				counter = v
				mu.Unlock()
				wg.Done()
			}()

		}

		wg.Wait()
	fmt.Println("Goroutines:",runtime.NumGoroutine())
		fmt.Println(counter)

}


func waitingGroup(){
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	/*
		Mutual exclusion lock
	*/
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar ", i)
	}
	wg.Done()
}
