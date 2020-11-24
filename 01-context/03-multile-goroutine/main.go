package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	data := make(chan int, 100)

	go func(data chan<- int) {
		for i := 0; i < 100; i++ {
			data <- i
		}
	}(data)

	go worker(ctx, "node01", data)
	go worker(ctx, "node02", data)
	go worker(ctx, "node03", data)

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string, data <-chan int) {
	for {
		select {
		case c := <-data:
			if ctx.Err() != nil {
				fmt.Println(name, "stop to handle data:", ctx.Err().Error())
				return
			}
			fmt.Println(name, "got data value", c)
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println(name, "got the stop channel")
			return
		}
	}
}
