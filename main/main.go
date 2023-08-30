package main

import "fmt"

type Resource struct {
	ID int64
	aa subStr
}

type subStr struct {
	dd string
}

func main() {
	b := subStr{dd: "dasfh"}
	a := Resource{
		ID: 6,
		aa: b,
	}

	fmt.Println(a)
}
