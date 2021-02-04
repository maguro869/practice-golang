package main

import (
	"fmt"
	"myfunc"
)

func main(){
	fmt.Println("Hello world")
	myfunc.Hello()
	var fslice []int

	fslice = append(fslice, 1)
	fslice = append(fslice, 3)
	fmt.Printf("%v\n", fslice)

	numMap := make(map[string]int)

	numMap["one"] = 1
	numMap["two"] = 2

	fmt.Printf("numMap[\"one\"]%d\n", numMap["one"])
	fmt.Printf("numMap:%v\n", numMap)


	
}