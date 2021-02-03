package main // this code gives error on counter part cause it will run a number const whch will print 11111111111 10 times

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	//runtime.GOMAXPROCS(10)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go say()
		go hi()
	}
	wg.Wait()
}

func say() {
	m.RLock()
	fmt.Printf(" hello there %v\n", counter)
	m.RUnlock()
	wg.Done()
}
func hi() {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
