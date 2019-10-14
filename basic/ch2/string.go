package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// HasPrefix 判断s是否以prefix开头
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix 判断s是否以suffix结尾
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains 判断substr是否是s的子串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasPrefix("Hello, world", "Hello"))
	fmt.Println(HasSuffix("Hello, world", "world"))
	fmt.Println(Contains("Hello, go", "go"))

	s := "Hello, 世界"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	n := 0
	for range s {
		n++
	}
	fmt.Println(n, len(s), utf8.RuneCountInString(s))
	fmt.Println(strings.Fields("aaa bbb ccc \t ddd"))
}
