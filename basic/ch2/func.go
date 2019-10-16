// 位向量使用一个无符号整型的slice，每一位代表集合中的一个元素; 如果设置第i位的元素，则代表集合包含i。
package main

import (
	"bytes"
	"fmt"
)

// SystemBit 兼容32/64位系统
const SystemBit int = (32 << (^uint(0) >> 63))

// IntSet 表示位向量
type IntSet struct {
	words []uint
}

// Has 判断是否存在非负整数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/SystemBit, uint(x%SystemBit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加非负整数x
func (s *IntSet) Add(x int) {
	word, bit := x/SystemBit, uint(x%SystemBit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

// Len 长度(保留相同索引的不同位的内容)
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words { // 索引位内容
		for word != 0 { // 内索引位的内容
			count++
			word &= word - 1
		}
	}
	return count
}

// Remove 移除
func (s *IntSet) Remove(x int) {
	word, bit := x/SystemBit, uint(x%SystemBit)
	s.words[word] &^= 1 << bit // 与添加操作相反 通过&^(异或)
}

// Clear 清空
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0 // 仅将对应下标索引对应的内容置为0即可
	}
}

// UnionWith 对s和t做并集操作，结果保存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 以字符串{1 2 3}的形式输出s
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < SystemBit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", SystemBit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	fmt.Println(x.String())
}
