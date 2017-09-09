package main

import (
	"fmt"
	"unicode"
)

func reverse(in string) string {

	res := []rune(in)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[j], res[i] = res[i], res[j]
	}
	return string(res)

}
func doubleLength(in string) string {
	return in + in
}

func expand(in string, f func(string) string) string {

	fmt.Println("in:", in)
	//internal expand function
	exp := func(in string) string {

		var start bool = false
		var idx1, idx2 int
		buffer := ""
		prev := ""
		for i, c := range in {
			if c == '$' {
				if idx2 > idx1 {
					s := f(in[idx1 : idx2+1])
					buffer += (s)
				} else if prev == string(c) { //handle $ repeats
					buffer += prev
				} else if i == (len(in) - 1) { //last char in string is $
					buffer += string(c)
				}

				idx1 = i
				idx2 = i
				start = true
				//		fmt.Println("set true ", c, i, unicode.IsLetter(c))

			} else if start == true && unicode.IsLetter(c) {
				idx2 = i

			} else if start == true && !unicode.IsLetter(c) && (idx2 > idx1) {
				s := f(in[idx1 : idx2+1])
				buffer += (s + string(c))
				idx1 = idx2
				start = false
				//		fmt.Println("set false ", c, i, unicode.IsLetter(c))
			} else {
				buffer += string(c)
			}
			prev = string(c)

		}
		//update case where $x is last string
		if idx2 > idx1 && start == true {
			s := f(in[idx1 : idx2+1])
			buffer += s

		}
		return buffer
	}

	result := exp(in)

	return result
}

func main() {

	r := expand("The $little big said $car park$H$A$$$ABC $", reverse)
	fmt.Println("out:", r)

	r = expand("The $little big said $car park$H$A$$$ABC $", doubleLength)

	fmt.Println("out:", r)

	r = expand("$The $little big said $car park$H$A$$$ABC $", reverse)
	fmt.Println("out:", r)
}
