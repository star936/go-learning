package v2_test

import (
	"testing"
	"work/basic/ch3/memo"
	v2 "work/basic/ch3/v2"
)

func TestMemo(t *testing.T) {
	m := v2.New(memo.HTTPGetBody)
	defer m.Close()
	memo.Concurrent(t, m)
}
