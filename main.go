package main

import "fmt"

type Hoge struct {
	N int
}

type Fuga struct {
	Hoge
}

func main() {
	f := Fuga{Hoge{N: 100}}
	fmt.Println(f.N)
	fmt.Println(f.Hoge.N)
}
