package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// strings.Split(対称文字列, 区切り文字列)
func main() {
	x := input("type strings")
	ar := strings.Split(x, " ")
	fmt.Println(ar)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
