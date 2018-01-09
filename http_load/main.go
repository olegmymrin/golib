package main

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func run(count int, out chan time.Duration, wg *sync.WaitGroup) {
	for i:= 0; i < count; i++ {
		start := time.Now()
		http.Get("http://127.0.0.1:30700/")
		elasped := time.Since(start)
		out<- elasped
	}
	wg.Done()
}

func read(in chan time.Duration) []time.Duration {
	total := time.Duration(0)
	min := time.Duration(0)
	max := time.Duration(0)
	var elapsed time.Duration
	ok := true
	for ok {
		elapsed, ok = <-in
		total += elapsed
		if elapsed < min {
			min = elapsed
		}
		if elapsed > max {
			max = elapsed
		}
	}
	return []time.Duration{total, min, max}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("specify count")
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	data := make(chan time.Duration, count)
	wg := sync.WaitGroup{}
	clients := 10
	for i:= 0; i < clients; i++{
		wg.Add(1)
		go run(count/clients, data, &wg)
	}
	wg.Wait()
	close(data)
	res := read(data)
	avg := res[0] / time.Duration(count);
	fmt.Println("count", count)
	fmt.Println("total", res[0].String())
	fmt.Println("avg", avg.String())
	fmt.Println("min", res[1].String())
	fmt.Println("max", res[2].String())
}
