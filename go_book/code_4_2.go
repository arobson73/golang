package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	fmt.Printf("reverse\n")
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//reverse
	reverse(a[:])
	fmt.Println(a)
	reverse(a[:])
	//rotate 1
	fmt.Printf("rotate\n")
	rotateN(a[:], 2)
	fmt.Println(a)
	//reverse pointer
	fmt.Printf("reverse pointer\n")
	reverseP(&a)
	fmt.Println(a)
	//rotate single pass
	fmt.Printf("rotate single pass\n")
	rotateSinglePass(a[:])
	fmt.Println(a)
	//
	fmt.Printf("remove string dupes\n")
	b := []string{"h", "h", "h", "e", "e", "e", "l", "l", "o", "o", "o"}
	fmt.Println(b)
	fmt.Println(removeStrDups(b))

	//create utf8 string
	fmt.Println("reverse utf string")
	u := []byte("Räksmörgås")
	fmt.Println(u)
	reverseRune(u)
	s := bytes2Str(u)
	fmt.Println(s)

	//squash space
	fmt.Println("squash space")
	sq := []byte("\t\n \t pppp   3r\n\n\n4")
	fmt.Println(sq)
	o := squashSpace(sq)
	fmt.Println(o)
	fmt.Println(bytes2Str(o))

}

//this is slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateN(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateSinglePass(s []int) {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first
}

func reverseP(s *[8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func removeStrDups(s []string) []string {
	idx := 0
	for _, c := range s {
		//	fmt.Println(ii, c, s[idx])
		if s[idx] == c {
			continue
		}
		idx++
		s[idx] = c
	}
	//fmt.Println(idx)
	return s[:idx+1]
}

//reverse bytes
func reverseB(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//reverse each rune then the whole thing
func reverseRune(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:]) //decodes first run
		reverseB(b[i : i+size])
		i += size
	}
	reverseB(b)
}

func bytes2Str(b []byte) string {
	n := len(b)
	fmt.Println(n)
	return string(b[:n])
}

func squashSpace(b []byte) []byte {
	out := b[:0]
	flag := false
	for _, c := range b {
		if unicode.IsSpace(rune(c)) {
			if !flag {
				out = append(out, ' ')
				flag = true
			} else {
				continue
			}
		} else {
			out = append(out, c)
			flag = false
		}
	}
	return out
}
