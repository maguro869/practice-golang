package main

import (
	"fmt"
	"strconv"
)

// Element 空のinterface
type Element interface{}

// List Element型
type List [] Element

// Person struct
type Person struct {
	name string
	age int
}

// Stringメソッドを定義 fmt.Stringerを実装
func (p Person) String() string {
	return "(name: "+ p.name + " - age: "+strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1 	// これはint
	list[1] = "Hello"	// これはstring
	list[2] = Person{"Dennis", 70}	// これはPerson

	// "Comma-okアサーション"で、型を判断する
	fmt.Println("------------Comma-okアサーション------------")
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

	fmt.Println()

	// "switchテスト"で、型を判断する
	fmt.Println("----------------switchテスト----------------")
	for index, element := range list {
		switch value := element.(type) {
			case int:
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string:
				fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
			case Person:
				fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
			default:
				fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}