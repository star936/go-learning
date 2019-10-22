package v1_test

import (
	"testing"
	"work/basic/ch3/memo"
	v1 "work/basic/ch3/v1"
)

func TestMemo(t *testing.T) {
	m := v1.New(memo.HTTPGetBody)
	memo.Concurrent(t, m)
}
