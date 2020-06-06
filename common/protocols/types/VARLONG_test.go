package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldVarlong(t *testing.T) {
	cut, err := types.NewFieldVarlong("TestVarlong", "int32 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestVarlong" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldVarlongWithDefault(t *testing.T) {
	cut, err := types.NewFieldVarlongWithDefault(
		"TestVarlong",
		"varlong test doc",
		32,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestVarlong" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int64(32) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypVARLONG_WriteAndRead(t *testing.T) {
	cut := types.TypVARLONG{}
	buf := &bytes.Buffer{}
	for _, i := range []int64{
		-64, -1, 0, 1, 63,
		-2147483648,
		2147483647,
		-9223372036854775808,
		9223372036854775807,
	} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("i: %v, err: %v", i, err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("i: %v, r: %v, err: %v", i, r, err)
		}
		if v, ok := r.(int64); !ok {
			t.Fatalf("i: %v, v: %v", i, v)
		} else if v != i {
			t.Fatalf("i: %v, v: %v", i, v)
		}
		buf.Reset()
	}
}
