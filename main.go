package main

type escape struct{}

func f() {
	g()
}

func g() {
	panic(escape{})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(escape); ok {
				println("Escaped")
			} else {
				panic(r)
			}
		}
	}()
	f()
}
