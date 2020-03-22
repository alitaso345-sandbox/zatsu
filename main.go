package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("別のゴールーチン")
	}()
	fmt.Println("mainゴールーチン")
	time.Sleep(50 * time.Millisecond)
}
