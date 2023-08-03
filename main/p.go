package main

import (
	"GolandProjects/addSum"
	"fmt"
)

func main() {
	addSum.AddFunction(1, 2)
	var x int64 = 1
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
}
