package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(j int) {
			wg.Add(1)
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
			fmt.Println("Routine", j, "is done.")
		}(i)
	}
	fmt.Println("Main routine is blocked")
	// wg.Wait()
	fmt.Println("Main routine is done")
}
