package main

import "fmt"

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func max2(val1 int, valsn ...int) int {
	if len(valsn) == 0 {
		return val1
	}
	max := val1
	for _, val := range valsn {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func min2(val1 int, valsn ...int) int {
	if len(valsn) == 0 {
		return val1
	}
	min := val1
	for _, val := range valsn {
		if val < min {
			min = val
		}
	}
	return min
}

//!-

func main() {
	fmt.Println("test max")
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(max())

	fmt.Println(max(4))
	fmt.Println(max(-100, 4, 3000))

	fmt.Println("test min")
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(min())

	fmt.Println(min(4))
	fmt.Println(min(-100, 4, 3000))

	fmt.Println("test max2")
	fmt.Println(max2(44, 5, 6, 2, 333, 5))

	fmt.Println("test min2")
	fmt.Println(min2(44, 5, 6, 2, 333, 5))

}
