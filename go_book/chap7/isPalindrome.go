package main

import (
	"fmt"
	"sort"
)

type palindrome []byte

func (x palindrome) Len() int           { return len(x) }
func (x palindrome) Less(i, j int) bool { return x[i] < x[j] }
func (x palindrome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		r := !s.Less(i, j) && !s.Less(j, i)
		if !r { //note check both index since Less is just < so can't have ABCDEF passing hence do both sides
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(palindrome([]byte("123456"))))
	fmt.Println(IsPalindrome(palindrome([]byte("123321"))))
	fmt.Println(IsPalindrome(palindrome([]byte("1234321"))))
}
