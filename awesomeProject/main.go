package main

import (
	"awesomeProject/singleton"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		go func ()  {
			fmt.Printf("%p\n", singleton.GetInstance())
		} ()
	}
}
