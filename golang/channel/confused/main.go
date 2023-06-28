package main

import (
	"fmt"
)

// run wrong with expect, all process run in first goroutine
func main() {
	var rdNums []int
	for i := 0; i < 1000; i++ {
		rdNums = append(rdNums, i)
	}

	pl := generatePipeline(rdNums)
	c1 := fanOut(pl, "1")
	c2 := fanOut(pl, "2")
	c3 := fanOut(pl, "3")
	c4 := fanOut(pl, "4")

	c := fanIn(c1, c2, c3, c4)
	sum := 0
	for i := range c {
		sum += i
	}

	fmt.Println(sum)
}

func generatePipeline(arrNums []int) <-chan int {
	pl := make(chan int, 100)
	go func() {
		for _, n := range arrNums {
			pl <- n
		}

		close(pl)
	}()

	return pl
}

func fanOut(in <-chan int, name string) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			fmt.Printf("Push square of %d to channel %s \n", v*v, name)
			out <- v * v

		}

		close(out)
	}()

	return out
}

func fanIn(inputChan ...<-chan int) <-chan int {
	in := make(chan int)

	go func() {
		for _, c := range inputChan {
			for v := range c {
				in <- v
			}
		}

		close(in)
	}()

	return in
}
