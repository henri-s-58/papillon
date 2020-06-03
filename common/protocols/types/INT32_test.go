package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldInt32(t *testing.T) {
	cut, err := types.NewFieldInt32("TestInt32", "int32 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt32" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldInt32WithDefault(t *testing.T) {
	cut, err := types.NewFieldInt32WithDefault(
		"TestInt32",
		"int32 test doc",
		-32,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt32" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int32(-32) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypINT32_WriteAndRead(t *testing.T) {
	cut := types.TypINT32{}
	buf := &bytes.Buffer{}
	for _, i := range []int32{-2147483648, -1, 0, 1, 2147483647} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(int32); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
