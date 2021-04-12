package main

import (
	"sync"
	"time"
)

var controller sync.Mutex
var counter = 0

func main() {
	go func() { controller.Lock() }()
	time.Sleep(time.Millisecond * 10)
	controller.Lock()
}

