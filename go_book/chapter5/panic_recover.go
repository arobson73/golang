package main

import (
	"fmt"
)

func math(a, b int) {
	type divzero struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		//no panic

		case divzero{}:
			fmt.Println("Divide by zero encountered")

		default:
			fmt.Println("%v ", p)
		}
	}()
	//alternative have no checking just print a/b, the recover takes care
	if b != 0 {
		fmt.Printf("res = %d\n", a/b)
	} else {
		panic(divzero{})
	}

}

func main() {
	math(10, 2)
	math(10, 0)

	fmt.Println("Hello something else")
}
