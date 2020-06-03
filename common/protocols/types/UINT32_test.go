package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldUint32(t *testing.T) {
	cut, err := types.NewFieldUint32("TestUint32", "Uint32 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestUint32" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldUint32WithDefault(t *testing.T) {
	cut, err := types.NewFieldInt32WithDefault(
		"TestUint32",
		"uint32 test doc",
		32,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestUint32" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int32(32) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypUINT32_WriteAndRead(t *testing.T) {
	cut := types.TypUINT32{}
	buf := &bytes.Buffer{}
	for _, i := range []uint32{0, 1, 2147483647, 4294967295} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(uint32); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
