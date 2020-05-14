package main

import "fmt"

func main() {
	one  := []string{"1","2",}
	two := &one  //引用
	(*two)[0] = "100"
	fmt.Println(one)
	fmt.Println(*two)
}
