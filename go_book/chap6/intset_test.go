// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Test_Length2(t *testing.T) {
	var x IntSet
	x.Add(2)
	x.Add(4)
	l := x.Len()
	if l != 2 {
		t.Error("For", x.String(), "expected", 2, "got", l)
	}
}

func Test_Remove(t *testing.T) {
	var x IntSet
	x.Add(9)
	x.Add(10)
	fmt.Println("Testing Remove")
	fmt.Println(&x) // "{9 10}
	x.Remove(10)
	fmt.Println(&x) // "{9}"

}

func Test_Clear(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(50)
	x.Add(333)
	fmt.Println("Testing Clear")
	fmt.Println(&x) // "{50 100 333}"
	x.Clear()
	fmt.Println(&x) // "{}"

}

func Test_Copy(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(333)
	fmt.Println("Testing Copy")
	fmt.Println(&x) // "{100 200 333}"
	y := x.Copy()   //this is a pointer, so in no need for &
	fmt.Println(y)  // "{100 200 333}"
}

func Test_AddAll(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	fmt.Println("Testing AddAll")
	fmt.Println(&x) // "{100 200 333}"
	x.AddAll(2, 4, 6)
	fmt.Println(&x) // "{2 4 6 100 200 333}"

}

func Test_Union(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	x.Add(500)
	fmt.Println("Testing Union")
	fmt.Println(&x)
	var y IntSet
	y.Add(222)
	y.Add(333)
	y.Add(444)
	fmt.Println(&y)
	x.UnionWith(&y)
	fmt.Println(&x)

}

func Test_IntersectWith(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	fmt.Println("Testing IntersectWith")
	fmt.Println(&x) // "{100 200 300}"
	fmt.Println(&y) // "{200 333 600}"
	x.IntersectWith(&y)
	fmt.Println(&x) // "{200}"

}
func Test_IntersectWith2(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	x.Add(400)
	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	fmt.Println("Testing IntersectWith2")
	fmt.Println(&x) // "{100 200 300 400}"
	fmt.Println(&y) // "{200 333 600}"
	x.IntersectWith(&y)
	fmt.Println(&x) // "{200}"

}

func Test_DifferenceWith(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	fmt.Println("Testing DifferenceWith")
	fmt.Println(&x) // "{100 200 300}"
	fmt.Println(&y) // "{200 333 600}"
	x.DifferenceWith(&y)
	fmt.Println(&x) // "{100 300}"

}

//try larger T set
func Test_DifferenceWith2(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	y.Add(400)
	fmt.Println("Testing DifferenceWith2")
	fmt.Println(&x) // "{100 200 300}"
	fmt.Println(&y) // "{200 333 400 600}"
	x.DifferenceWith(&y)
	fmt.Println(&x) // "{100 300}"

}

//try larger T set
func Test_DifferenceWith3(t *testing.T) {
	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	x.Add(400)
	x.Add(500)
	x.Add(5000)

	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	fmt.Println("Testing DifferenceWith3")
	fmt.Println(&x) // "{100 200 300 400 500 5000}"
	fmt.Println(&y) // "{200 333 600}"
	x.DifferenceWith(&y)
	fmt.Println(&x) // "{100 300 400 500 5000}"

}

func Test_SymmetricDifference(t *testing.T) {

	var x IntSet
	x.Add(100)
	x.Add(200)
	x.Add(300)
	x.Add(400)
	var y IntSet
	y.Add(200)
	y.Add(600)
	y.Add(333)
	y.Add(700)
	y.Add(800)

	fmt.Println("Testing SymmetricDifference")
	fmt.Println(&x) // "{100 200 300 400}"
	fmt.Println(&y) // "{200 333 600 700}"
	x.SymetricDifference(&y)
	fmt.Println(&x) // "{100 300 400 U 600 333 700} => {100 300 333 400 600 700}"

}
