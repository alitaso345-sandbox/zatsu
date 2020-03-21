package main

import "fmt"

type Stringer interface {
	String() string
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type Octal int

func (o Octal) String() string {
	return fmt.Sprintf("%o", int(o))
}

type Binary int

func (b Binary) String() string {
	return fmt.Sprintf("%b", int(b))
}

type Func func() string

func f(s Stringer) {
	switch v := s.(type) {
	case Hex:
		fmt.Println(int(v), "Hex")
	case Octal:
		fmt.Println(int(v), "Octal")
	case Binary:
		fmt.Println(int(v), "Binary")
	}
}

func main() {
	var s1 Stringer = Hex(100)
	var s2 Stringer = Octal(100)
	var s3 Stringer = Binary(100)

	fmt.Println(s1.String())
	fmt.Println(s2.String())
	fmt.Println(s3.String())

	f(Hex(100))
	f(Octal(100))
	f(Binary(100))
}
