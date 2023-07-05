package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	strings := [5]string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	addr := uintptr(unsafe.Pointer(&strings))
	println("start here", addr)
	for i, fruit := range strings {
		addrEl := uintptr(unsafe.Pointer(&strings[i]))
		addrFor := uintptr(unsafe.Pointer(&fruit))
		println(i, addrFor, "value")
		println(i, addrEl, "pointer")

		headerEl := (*reflect.StringHeader)(unsafe.Pointer(&strings[i]))
		println("header address copy strings", headerEl.Data)

		header := (*reflect.StringHeader)(unsafe.Pointer(&fruit))
		println("header address origin strings", header.Data)

		for i := 0; i < len(strings[i]); i++ {
			ptr := (*byte)(unsafe.Pointer(header.Data + uintptr(i)))
			fmt.Printf(",%c ", *ptr)
		}
		println()

	}
}
