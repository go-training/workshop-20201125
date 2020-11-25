package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	outChan := make(chan int)
	errChan := make(chan error)
	finishChan := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errChan chan<- error, val int) {
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			outChan <- val
			defer wg.Done()
		}(outChan, errChan, i)
	}

	go func() {
		wg.Wait()
		close(finishChan)
	}()

Loop:
	for {
		select {
		case v := <-outChan:
			log.Println(v)
		case v := <-errChan:
			log.Println(v)
		case <-time.After(5 * time.Second):
			log.Println("timeout")
			break Loop
		case <-finishChan:
			log.Println("finished all jobs")
			break Loop
		}
	}

	log.Println("finish main")
}
