package types_test

import (
	"bytes"
	"math"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldFloat64(t *testing.T) {
	cut, err := types.NewFieldFloat64("TestFloat64", "int64 test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestFloat64" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldFloat64WithDefault(t *testing.T) {
	cut, err := types.NewFieldFloat64WithDefault(
		"TestFloat64",
		"int64 test doc",
		-64,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestFloat64" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if int64(-64) != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypFLOAT64_WriteAndRead(t *testing.T) {
	cut := types.TypFLOAT64{}
	buf := &bytes.Buffer{}
	for _, i := range []float64{-9.0, -1.0, 0.0, 1.0, 9.0, math.MaxFloat64} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(float64); !ok {
			t.Fatalf("no float64. v: %v, i: %v", v, i)
		} else if v != i {
			t.Fatalf("v: %v, i: %v", v, i)
		}
		buf.Reset()
	}
}
