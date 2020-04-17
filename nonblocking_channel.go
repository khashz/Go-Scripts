package main

import (
	"fmt"
	"time"
	"sync"
)

func executor(execorname string, speed time.Duration, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("Executor %s processed job %s \n", execorname, job)
		time.Sleep(speed * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	// If we allocate storage to the channel, it doesnt crash if no worker is
	// there to receive the message
	ch := make(chan string, 20)

	for i := 0; i < 20; i++ {
		jobname := fmt.Sprintf("%d", i)
		ch <- jobname
	}

	wg.Add(1)
	go executor("A", 150, ch, &wg)

	wg.Add(1)
	go executor("B", 500, ch, &wg)

	wg.Add(1)
	go executor("C", 250, ch, &wg)

	close(ch)

	wg.Wait()
}
