package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldInt16(t *testing.T) {
	cut, err := types.NewFieldInt16("TestInt16", "int16 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt16" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldInt16WithDefault(t *testing.T) {
	cut, err := types.NewFieldInt16WithDefault(
		"TestInt16",
		"int16 test doc",
		-16,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestInt16" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int16(-16) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypINT16_WriteAndRead(t *testing.T) {
	cut := types.TypINT16{}
	buf := &bytes.Buffer{}
	for _, i := range []int16{-32768, -1, 0, 1, 32767} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(int16); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
