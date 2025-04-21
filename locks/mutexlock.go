package main

import (
	"fmt"
	"sync"
)

func increment() {
	var mu sync.Mutex
	var counter int

	mu.Lock()
	counter++
	defer mu.Unlock()
	fmt.Println("count", counter)
}
