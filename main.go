package main

import "fmt"

type Stringer interface {
	String() string
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	var s Stringer = Hex(100)
	fmt.Println(s.String())
}
