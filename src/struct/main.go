package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main(){
	var tom person

	tom.name, tom.age = "Tom", 23

	bob := person{age:25, name:"Bob"}

	paul := person{"Paul", 25}

	fmt.Printf("%v\n", tom)
	fmt.Printf("%v\n", bob)
	fmt.Printf("%v\n", paul)
	
}