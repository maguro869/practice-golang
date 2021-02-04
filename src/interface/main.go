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
	loan float32
}
// Employee struct
type Employee struct {
	Human //匿名フィールド
	company string
	money float32
}

// SayHi method
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// SayHi Employee edition
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Sing method
func (h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la:", lyrics)
}

// Men interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}


func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-YYYY"}, "Things Ltd.", 5000}

	// Men型interfaceの変数
	var i Men

	i = mark
	fmt.Println("This is Mark, a Student:")
	i.SayHi()
	i.Sing("November rain")

	i = tom
	fmt.Println("This is Tom, an Employee")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)

	x[0], x[1], x[2] = mark, paul, sam

	for _, value := range x {
		value.SayHi()
	}

}