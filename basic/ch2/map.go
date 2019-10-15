package main

import (
	"fmt"
)

// equal 判断map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	m := map[string]int{"A": 0}
	n := map[string]int{"B": 42}
	fmt.Println(equal(m, n))
}
