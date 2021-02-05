package main

import (
	"fmt"
)

func main(){
	var a int = 1
	var b = 2
	i := 1

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", i)

	name := "maguro"
	desc := `
	He is Tsuna.
	He like Tsuna.
	`

	fmt.Printf("%s\n", name)
	fmt.Printf("%c\n", name[0])
	fmt.Printf("%s\n", desc)

	arr := [3]int{}
	arr[0] = 1
	fmt.Printf("%d\n",arr[0])
	fmt.Printf("%d\n", arr[1])




	
}