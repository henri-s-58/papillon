package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldVarint(t *testing.T) {
	cut, err := types.NewFieldVarint("TestVarint", "int32 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestVarint" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldVarintWithDefault(t *testing.T) {
	cut, err := types.NewFieldVarintWithDefault(
		"TestVarint",
		"varint test doc",
		32,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestVarint" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if 32 != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypVARINT_WriteAndRead(t *testing.T) {
	cut := types.TypVARINT{}
	buf := &bytes.Buffer{}
	for _, i := range []int{-64, -1, 0, 1, 63, -2147483648, 2147483647} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("i: %v, err: %v", i, err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("i: %v, err: %v", i, err)
		}
		if v, ok := r.(int); !ok {
			t.Fatalf("i: %v, v: %v", i, v)
		} else if v != i {
			t.Fatalf("i: %v, v: %v", i, v)
		}
		buf.Reset()
	}
}
