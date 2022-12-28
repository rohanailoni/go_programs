package main

import (
	"fmt"
	"sync"
	"time"
)

/*
this code demonstrates starvation in go where the geedy is holding
the lock form the non greedy function and doing trated unfairly
*/
func main() {
	var waitgroup sync.WaitGroup
	var lock sync.Mutex
	greedy := func() {
		defer waitgroup.Done()

		var count int = 0
		for begin := time.Now(); time.Since(begin) <= 2*time.Second; {
			lock.Lock()
			time.Sleep(3 * time.Second)
			lock.Unlock()

			count++
		}
		fmt.Println("the greeedy one ran around", count, " times")

	}
	nongreddy := func() {
		defer waitgroup.Done()

		var count int = 0
		for begin := time.Now(); time.Since(begin) <= 2*time.Second; {
			lock.Lock()
			time.Sleep(1 * time.Second)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Second)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Second)
			lock.Unlock()
			count++
		}
		fmt.Println("the non greeedy one ran around", count, " times")
	}
	waitgroup.Add(2)
	go greedy()
	go nongreddy()
	waitgroup.Wait()

}
