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

	// Since our channel is empty, we have to have listeners ready to remove
	// messages from our channel
	ch := make(chan string)

	wg.Add(1)
	go executor("A", 150, ch, &wg)

	wg.Add(1)
	go executor("B", 500, ch, &wg)

	wg.Add(1)
	go executor("C", 250, ch, &wg)

	for i := 0; i < 10; i++ {
		jobname := fmt.Sprintf("%d", i)
		ch <- jobname
	}

	close(ch)

	wg.Wait()
}
