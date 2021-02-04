package main

import (
	"fmt"
)

// Human struct
type Human struct {
	name string
	age int
	phone string
}
// Student struct
type Student struct {
	Human //匿名フィールド
	school string
}
// Employee struct
type Employee struct {
	Human //匿名フィールド
	company string
}

// SayHi method
func (h * Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// SayHi Employee edition
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}



func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}