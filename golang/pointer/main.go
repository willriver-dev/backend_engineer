package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"dev": "thang ha", "test": "huong nguyen"}
	b := a
	fmt.Println(a, b)
	a["dev"] = "thuong bk"
	fmt.Println(a, b)
	// map and slice is pointer (refrerence)
}
