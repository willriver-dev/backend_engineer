package main

import (
	"backend_engineer/singleton_dp/singleton"
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func ()  {
			fmt.Printf("%p\n", singleton.GetInstance())
		} ()
	}
	time.Sleep(time.Second*5)
}
