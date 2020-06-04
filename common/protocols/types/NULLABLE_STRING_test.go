package types_test

import (
	"bytes"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldNullableString(t *testing.T) {
	cut, err := types.NewFieldNullableString("TestNullableString", "Nullable string test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestNullableString" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldNullableStringWithDefault(t *testing.T) {
	cut, err := types.NewFieldNullableStringWithDefault(
		"TestNullableString",
		"Nullable string test doc",
		"hello Nullable string",
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestNullableString" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if "hello Nullable string" != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypNULLABLE_STRING_WriteAndRead(t *testing.T) {
	cut := types.TypNULLABLE_STRING{}
	buf := &bytes.Buffer{}
	for _, i := range []interface{}{nil, "Hello", "세계", "!"} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(string); !ok {
			if r != nil {
				t.Fatalf("v: %v", v)
			}
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
