package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	ch := input(os.Stdin)
	problems := []string{
		"hello",
		"world",
		"alice",
		"bob",
		"cathy",
		"domain",
		"easy",
	}

	for i := 0; i < len(problems); i++ {
		fmt.Print(problems[i] + "> ")

		select {
		case v := <-ch:
			fmt.Println(">> " + v)
		case <-time.After(5 * time.Second):
			fmt.Println("time up")
			i = len(problems)
		}
	}
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}
