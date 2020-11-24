package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working on job", strconv.Itoa(i))
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}
