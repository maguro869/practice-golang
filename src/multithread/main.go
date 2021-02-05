package main

import (
	"fmt"
	"runtime"
) 

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()	// 別のgorutineにCPU使用時間を与える
		fmt.Println(s)
	}
}

func main() {
	go say("world")	// 新しいGorutines
	say("hello")	// 今のGorutines
}