package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	stop := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-stop:
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
	close(stop)
	time.Sleep(5 * time.Second)
}
