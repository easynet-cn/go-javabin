package javabin

import (
	"bytes"
	"encoding/binary"

	"github.com/google/uuid"
)

func UUIDBytes(uuid uuid.UUID, serialVersionUID int64) []byte {
	byteBuffer := new(bytes.Buffer)

	byteBuffer.Write(ShortBytes(STREAM_MAGIC))
	byteBuffer.Write(ShortBytes(STREAM_VERSION))
	byteBuffer.Write([]byte{TC_OBJECT})
	byteBuffer.Write([]byte{TC_CLASSDESC})

	className := "java.util.UUID"
	classNameLen := len(className)

	byteBuffer.Write(ShortBytes(uint16(classNameLen)))
	byteBuffer.Write([]byte(className))

	byteBuffer.Write(LongBytes(uint64(serialVersionUID)))

	//flags
	byteBuffer.Write([]byte{2})

	//fields length
	byteBuffer.Write(ShortBytes(2))

	//filed type code
	byteBuffer.Write([]byte{LONG_TYPE})

	f1 := "leastSigBits"
	f1Len := len(f1)

	byteBuffer.Write(ShortBytes(uint16(f1Len)))
	byteBuffer.Write([]byte(f1))

	//filed type code
	byteBuffer.Write([]byte{LONG_TYPE})

	f2 := "mostSigBits"
	f2Len := len(f2)

	byteBuffer.Write(ShortBytes(uint16(f2Len)))
	byteBuffer.Write([]byte(f2))

	byteBuffer.Write([]byte{TC_ENDBLOCKDATA})
	byteBuffer.Write([]byte{TC_NULL})

	mostSigBits, leastSigBits := UUIDSigBits(uuid)

	byteBuffer.Write(LongBytes(uint64(leastSigBits)))

	byteBuffer.Write(LongBytes(uint64(mostSigBits)))

	return byteBuffer.Bytes()

}

func UUIDSigBits(uuid uuid.UUID) (int64, int64) {
	msb := int64(0)
	lsb := int64(0)

	for i, b := range uuid {
		if i >= 0 && i < 8 {
			msb = (msb << 8) | int64(b&0xff)
		} else if i >= 8 && i < 16 {
			lsb = (lsb << 8) | int64(b&0xff)
		}
	}

	return msb, lsb

}

func StringBytes(str string) []byte {
	buf := new(bytes.Buffer)

	buf.Write(ShortBytes(STREAM_MAGIC))
	buf.Write(ShortBytes(STREAM_VERSION))
	buf.Write([]byte{TC_STRING})

	strLen := len(str)

	buf.Write(ShortBytes(uint16(strLen)))
	buf.Write([]byte(str))

	return buf.Bytes()
}

func ShortBytes(i uint16) []byte {
	bytes := make([]byte, 2)

	binary.BigEndian.PutUint16(bytes, i)

	return bytes
}

func LongBytes(i uint64) []byte {
	bytes := make([]byte, 8)

	binary.BigEndian.PutUint64(bytes, i)

	return bytes
}
