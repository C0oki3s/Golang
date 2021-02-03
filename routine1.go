package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() //this will read lock here instead in the func say()
		go say()
		m.Lock() // this will lock counter to complete one work and then move on in func hi()
		go hi()
	}
	wg.Wait()
}
func say() {

	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}
func hi() {
	counter++
	m.Unlock()
	wg.Done()
}

// package main    | ---> code without mutex

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var counter = 0

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
//
// 		go say()
//
// 		go hi()
// 	}
// 	wg.Wait()
// }
// func say() {

// 	fmt.Printf("Hello %v\n", counter)
//
// 	wg.Done()
// }
// func hi() {
// 	counter++
// 	wg.Done()
// }
