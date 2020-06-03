package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldInt8(t *testing.T) {
	cut, err := types.NewFieldInt8("TestInt8", "int8 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt8" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldInt8WithDefault(t *testing.T) {
	cut, err := types.NewFieldInt8WithDefault(
		"TestInt8",
		"int8 test doc",
		-1,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt8" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int8(-1) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypINT8_WriteAndRead(t *testing.T) {
	cut := types.TypINT8{}
	// -128 through 127.
	buf := &bytes.Buffer{}
	for _, i := range []int8{-128, -1, 0, 1, 127} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(int8); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
