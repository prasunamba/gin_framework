package main

import (
	"fmt"
	"sync"
)

var rwlock sync.RWMutex

var count int

func rwlockfunc() {
	write()
	read()
}
func read() {
	rwlock.RLock()
	fmt.Println("count", count)
	defer rwlock.RUnlock()
}
func write() {
	fmt.Println("write done")
	rwlock.Lock()
	defer rwlock.Unlock()
	count++
}
