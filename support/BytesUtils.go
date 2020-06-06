package support

import (
	"bytes"
	"papillon/errorx"
)

func SizeOfUnsignedVarint(value uint) int {
	bl := 1
	v := value
	for ; (v & 0xffffff80) != 0; {
		bl += 1
		v >>= 7
	}
	if bl < 1 {
		return 1
	}
	return bl
}

func SizeOfVarint(value int) int {
	return SizeOfUnsignedVarint(uint((value << 1) ^ (value >> 31)))
}

func SizeOfVarlong(value int64) int {
	var v = uint64((value << 1) ^ (value >> 63))
	bl := 1
	const b uint64 = 0xffffffffffffff80
	for ; (v & b) != uint64(0); {
		bl += 1
		v >>= 7
	}
	if bl < 1 {
		return 1
	}
	return bl
}

func WriteUnsignedVarint(value uint, buf *bytes.Buffer) {
	for ; (value & 0xffffff80) != 0; value >>= 7 {
		var b = (byte)((value & 0x7f) | 0x80)
		buf.WriteByte(b)
	}
	buf.WriteByte(byte(value))
}

func WriteVarint(value int, buf *bytes.Buffer) {
	u := uint((value << 1) ^ (value >> 31))
	WriteUnsignedVarint(u, buf)
}

func WriteVarlong(value int64, buf *bytes.Buffer) {
	var v = uint64((value << 1) ^ (value >> 63))
	for ; (v & 0xffffffffffffff80) != uint64(0); {
		var b = uint8((v & 0x7f) | 0x80)
		buf.WriteByte(b)
		v >>= 7
	}
	buf.WriteByte(uint8(v))
}

func ReadUnsignedVarint(buf *bytes.Buffer) (uint, errorx.IllegalArgumentError) {
	value := uint(0)
	i := 0
	var bs []byte
	for bs = buf.Next(1); (bs[0] & 0x80) != 0; {
		value |= (uint(bs[0]) & 0x7f) << i
		i += 7
		if i > 28 {
			return 0, errorx.NewIllegalArgumentErrorf(
				"Varint is too long, the most significant bit in the 5th byte is set, converted value: %X",
				value,
			)
		}
		b, _ := buf.ReadByte()
		bs[0] = b
	}
	value |= uint(bs[0]) << i
	return value, nil
}

func ReadVarint(buf *bytes.Buffer) (int, errorx.IllegalArgumentError) {
	value, err := ReadUnsignedVarint(buf)
	if err != nil {
		return -1, err
	}
	return int((value >> 1) ^ -(value & 1)), nil
}

func ReadVarlong(buf *bytes.Buffer) (int64, errorx.IllegalArgumentError) {
	var value uint64 = 0
	var i = 0
	var b uint8
	for b, _ = buf.ReadByte(); (b & 0x80) != 0; b, _ = buf.ReadByte() {
		value |= (uint64(b) & 0x7f) << uint64(i)
		i += 7
		if i > 63 {
			return -1, errorx.NewIllegalArgumentErrorf("%d", value)
		}
	}
	value |= uint64(b) << uint64(i)
	r := (value >> 1) ^ -(value & 1)
	return int64(r), nil
}
