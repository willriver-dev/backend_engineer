package main

import (
	"backend_engineer/builder_pattern/builder"
	"fmt"
)

func main()  {
	normalBuilder := builder.GetBuilder("normal")
	//iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Printf("Normal house door type: %v\n", normalHouse.GetDoorType())
	fmt.Printf("Normal house window type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal house numbers flow: %d\n", normalHouse.GetNumbersFloor())
}
