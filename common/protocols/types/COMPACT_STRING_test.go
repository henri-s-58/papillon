package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldCompactString(t *testing.T) {
	cut, err := types.NewFieldCompactString("TestString", "string test doc")
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

func TestNewFieldCompactStringWithDefault(t *testing.T) {
	cut, err := types.NewFieldCompactStringWithDefault(
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

func TestTypCOMPACT_STRING_WriteAndRead(t *testing.T) {
	cut := types.TypCOMPACT_STRING{}
	buf := &bytes.Buffer{}
	for _, i := range []string{"Hello", "세계", "!"} {
		buf.Reset()
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
	}
}
