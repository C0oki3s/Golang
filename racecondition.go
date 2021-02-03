package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "there"
	go func() {
		fmt.Println(msg)
	}()
	msg = "hi"
	time.Sleep(1000 * time.Millisecond)
}
