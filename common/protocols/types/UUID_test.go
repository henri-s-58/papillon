package types_test

import (
	"bytes"
	"github.com/google/uuid"
	"papillon/common/protocols/types"
	"testing"
)

func TestNewFieldUUID(t *testing.T) {
	cut, err := types.NewFieldUUID("TestUUID", "uuid test doc")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestUUID" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if cut.HasDefaultValue() {
		t.Fatalf("want: have not default value")
	}
}

func TestNewFieldUUIDWithDefault(t *testing.T) {
	uid := uuid.New()
	cut, err := types.NewFieldUUIDWithDefault(
		"TestUUID",
		"uuid test doc",
		uid,
	)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if cut.Name() != "TestUUID" {
		t.Fatalf("cut name: %v", cut.Name())
	}
	if !cut.HasDefaultValue() {
		t.Fatalf("want: has default value")
	}
	if uid != cut.DefaultValue() {
		t.Fatalf("DefaultValue: %v", cut.DefaultValue())
	}
}

func TestTypUUID_WriteAndRead(t *testing.T) {
	cut := types.TypUUID{}
	buf := &bytes.Buffer{}
	for _, i := range []uuid.UUID{uuid.New(), uuid.New()} {
		err := cut.Write(buf, i)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		r, err := cut.Read(buf)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		if v, ok := r.(uuid.UUID); !ok {
			t.Fatalf("v: %v", v)
		} else if v != i {
			t.Fatalf("v: %v", v)
		}
		buf.Reset()
	}
}
