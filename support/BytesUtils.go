package support

import (
	"bytes"
	"papillon/errorx"
)

func SizeOfUnsignedVarint(value int) int {
	bytes := 1
	v := uint(value)
	for ; (v & 0xffffff80) != 0; {
		bytes += 1
		v >>= 7
	}
	return bytes
}

func WriteUnsignedVarint(value uint, buf *bytes.Buffer) {
	for ; (value & 0xffffff80) != 0; value >>= 7 {
		var b = (byte)((value & 0x7f) | 0x80)
		buf.WriteByte(b)
	}
	buf.WriteByte(byte(value))
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
	}
	value |= uint(bs[0]) << i
	return value, nil
}
