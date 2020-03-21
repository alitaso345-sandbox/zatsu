package main

import (
	"errors"
	"fmt"
	"log"
)

func f() (rerr error) {
	defer func() {
		if r := recover(); r != nil {
			if err, isErr := r.(error); isErr {
				rerr = err
			} else {
				rerr = fmt.Errorf("else: %v", r)
			}
		}
	}()
	panic(errors.New("new error"))
	return nil
}

func main() {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}
