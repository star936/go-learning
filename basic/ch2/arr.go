package main

import (
	"fmt"
	"strings"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice x仍有增长空间，扩展其空间
		z = x[:zlen]
	} else {
		// slice x已无空间，为其分配一个新的底层数组
		//
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// nonempty 去除slice中的空白字符
func nonempty(str []string) []string {
	out := str[:0] // 引用原始slice的新的零长度的slice
	for _, v := range str {
		if strings.TrimSpace(v) != "" {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
