package support_test

import (
	"papillon/support"
	"testing"
)

func TestSizeOfUnsignedVarint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		j := support.SizeOfUnsignedVarint(i)
		t.Log("i:", i, j)
		if i < 128 {
			if j == 1 {
				continue
			}
			t.Fatal(i, j)
		} else {
			if j == 2 {
				continue
			}
			t.Fatal(i, j)
		}
	}
}
