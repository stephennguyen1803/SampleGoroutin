package functionclosure

import "fmt"

func fibonnaci() func() int {
	a, b := 0, 1
	return func() int {
		current := a
		a, b = b, a+b
		return current
	}
}

func RunFiBonacci() {
	n := 10
	f := fibonnaci()
	fmt.Println("Fibonacci sequence with N:=", n)
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
}
