package main

import (
	"fmt"
)

// Skills 特技
type Skills []string

// Human 人間
type Human struct {
	name string
	age int
}

// Student 学生
type Student struct {
	Human
	department string

	Skills
	int
} 



func main() {
	bob := Student{Human: Human{"Bob", 20}, department: "Music"}

	fmt.Printf("%v\n", bob)
	fmt.Printf("His name is %s\n", bob.name)

	bob.int = 3
	bob.Skills = []string{"DJ"}
	fmt.Printf("His Student No. is %d\n", bob.int)
	fmt.Printf("His Skills is %s\n", bob.Skills[0])

	bob.Skills = append(bob.Skills, "Sing")
	fmt.Printf("His Skills : %v\n", bob.Skills)

}