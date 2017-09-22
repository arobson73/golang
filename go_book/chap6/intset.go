// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// return the number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		//		fmt.Printf("here %d %b\n", i, word)
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}

	return count
}

//remove value
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	var mask uint64 = ^(1 << bit)
	s.words[word] &= mask
}

//clear the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

//copy
func (s *IntSet) Copy() *IntSet {
	n := &IntSet{}
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return n
}

//addAll
func (s *IntSet) AddAll(in ...int) {
	for _, val := range in {
		s.Add(val)
	}
}

//intersectWith
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

//DifferenceWith S\T => 1010 \ 1001 => 0010
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) && word != 0 {
			s.words[i] &^= word
		}
	}
}

//symmetric difference S\T U T\S
func (s *IntSet) SymetricDifference(t *IntSet) {

	c := t.Copy()       // copy contents of t
	t.DifferenceWith(s) // mutate t
	s.DifferenceWith(c) // mutate s
	s.UnionWith(t)      //union

}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
