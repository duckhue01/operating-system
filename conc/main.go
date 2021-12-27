package main

import (
	"sync"
	"fmt"
	// "time"
)

var mutex1 = sync.Mutex{}
var mutex2 = sync.Mutex{}
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go thread2()
	go thread1()

	wg.Wait()

}
func thread1() {
	mutex1.Lock()
	for i := 1; i < 100000000;i++ {

	}
	mutex2.Lock()
	fmt.Println("thread 1")
	mutex1.Unlock()
	mutex2.Unlock()
	wg.Done()
}

func thread2() {
	mutex2.Lock()
	// for i := 1; i < 10000000000;i++ {

	// }
	mutex1.Lock()
	fmt.Println("thread 2")
	mutex1.Unlock()
	mutex2.Unlock()
	wg.Done()

}
