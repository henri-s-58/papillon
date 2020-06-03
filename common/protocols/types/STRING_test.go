package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldString(t *testing.T) {
	cut, err := types.NewFieldString("TestString", "string test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestString" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldStringWithDefault(t *testing.T) {
	cut, err := types.NewFieldStringWithDefault(
		"TestString",
		"string test doc",
		"hello string",
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestString" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if "hello string" != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypSTRING_WriteAndRead(t *testing.T) {
	cut := types.TypSTRING{}
	buf := &bytes.Buffer{}
	for _, i := range []string{"Hello", "세계", "!"} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(string); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
