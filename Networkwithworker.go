package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, host string, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
func main() {
	my := flag.String("host", "127.0.0.1", "Enter A IP or Host(Ex:- -host='google.com' or -host=192.168.1.10)")
	name := flag.Int("Num", 1000, "Enter Nummber of Ports")
	flag.Parse()
	host := *my
	ports := make(chan int, 1000)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, host, results)
	}
	go func() {
		for i := 1; i <= *name; i++ {
			ports <- i
		}
	}()
	for i := 0; i < *name; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
