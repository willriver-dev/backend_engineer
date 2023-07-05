package main

import (
	"fmt"
	"time"
)

// run exactly, all task equally divided among 5 goroutine
func main_draft() {
	var rdNums []int
	for i := 0; i < 1000; i++ {
		rdNums = append(rdNums, i)
	}

	pl := createQueue(rdNums)
	for i := 0; i < 5; i++ {
		go process(pl, fmt.Sprintf("worker%d", i))
	}

	time.Sleep(1 * time.Minute)
}

func createQueue(arrNums []int) <-chan int {
	pl := make(chan int)
	go func() {
		for _, n := range arrNums {
			pl <- n
		}

		close(pl)
	}()

	return pl
}

func process(in <-chan int, name string) {
	count := 0
	go func() {
		for v := range in {
			fmt.Printf("Push square of %d to goroutine_channel %s \n", v*v, name)
			count++

		}

		fmt.Printf("Process %s success, total number reiceive %d\n", name, count)
	}()
}
