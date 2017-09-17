package main

import "fmt"

func joinStrings(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var res string
	for _, s := range strs {
		res += (s + sep)
	}
	return res
}

func main() {
	fmt.Println(joinStrings(" ", "hello", "how", "you", "doing"))

}
