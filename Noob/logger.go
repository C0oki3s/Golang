package main

import (
	"fmt"
	"time"
)

const (
	logInfo  = "INFO"
	logwarn  = "Warning"
	logError = "Error"
)

type log struct {
	time      time.Time
	serverity string
	msg       string
}

var ch = make(chan log, 50)

func main() {
	go logger()
	ch <- log{time.Now(), logInfo, "App is starting"}
	ch <- log{time.Now(), logInfo, "App is Ending"}
	time.Sleep(1000 * time.Millisecond)
}
func logger() {
	for entry := range ch {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.serverity, entry.msg)
	}
}
