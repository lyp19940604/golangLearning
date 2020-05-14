package main

import (
	"fmt"
	"regexp"
)

const text = `My email is lll@gmail.com
My email is wwww@gmail.com 
My email is qqq@gmail.com.cn `
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	match1 := re.FindAllStringSubmatchIndex(text,-1)
	fmt.Println(match)
	fmt.Println(match1)
}
