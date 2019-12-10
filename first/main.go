package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Gopher!") ; f.Println("hello again" + "!")
	if 5 > 1 {
		fmt.Println("Bigger")
	}
	fmt.Println(runtime.NumCPU() + 1)
}


