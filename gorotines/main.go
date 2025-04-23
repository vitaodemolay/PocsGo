package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sync sync.WaitGroup
	for i := 10; i > 0; i-- {
		sync.Add(1)
		go PrintNumerWithDelay(i, &sync)
	}
	sync.Wait()
	fmt.Println("All done!")
}

func PrintNumerWithDelay(number int, sync *sync.WaitGroup) {
	//time.Sleep(time.Microsecond * time.Duration(number) * 100)
	time.Sleep(time.Second * 1)
	fmt.Println("Done! Numer is ", number)
	sync.Done()
}
