package deserializer

import (
	"bytes"
	"encoding/binary"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"io"
)

func NewClassFileReader(data []byte) intsrv.IClassFileReader {
	return &ClassFileReader{
		reader: bytes.NewReader(data),
		data:   data,
	}
}

type ClassFileReader struct {
	reader *bytes.Reader
	data   []byte
}

func (r *ClassFileReader) Offset() int {
	currentOffset, _ := r.reader.Seek(0, io.SeekCurrent)
	return int(currentOffset)
}

func (r *ClassFileReader) Skip(length int) {
	_, _ = r.reader.Seek(int64(length), io.SeekCurrent)
}

func (r *ClassFileReader) Read() byte {
	b, _ := r.reader.ReadByte()
	return b
}

func (r *ClassFileReader) ReadUnsignedByte() int {
	var b uint8
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return int(b)
}

func (r *ClassFileReader) ReadUnsignedShort() int {
	var b uint16
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return int(b)
}

func (r *ClassFileReader) ReadMagic() intsrv.Magic {
	var b uint32
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return intsrv.Magic(b)
}

func (r *ClassFileReader) ReadInt() int {
	var b uint32
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return int(b)
}

func (r *ClassFileReader) ReadFloat() float32 {
	var b float32
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return b
}

func (r *ClassFileReader) ReadLong() int64 {
	var b int64
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return b
}

func (r *ClassFileReader) ReadDouble() float64 {
	var b float64
	_ = binary.Read(r.reader, binary.BigEndian, &b)
	return b
}

func (r *ClassFileReader) ReadFully(length int) []byte {
	ret := r.data[r.Offset() : r.Offset()+length]
	r.Skip(length)
	return ret
}

func (r *ClassFileReader) ReadUTF8() string {
	utflenx := r.ReadUnsignedShort()
	maxOffset := r.Offset() + utflenx

	charArray := make([]rune, utflenx)
	charArrayOffset := 0

	for r.Offset() < maxOffset {
		c := r.Read()

		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			// 0xxxxxxx
			charArray[charArrayOffset] = rune(c)
			charArrayOffset++
		case 12, 13:
			if r.Offset()+1 > maxOffset {
				return ""
			}
			char2 := r.Read()
			if (char2 & 0xC0) != 0x80 {
				return ""
			}
			charArray[charArrayOffset] = rune(((c & 0x1F) << 6) | (char2 & 0x3F))
			charArrayOffset++
		case 14:
			// 1110xxxx 10xxxxxx 10xxxxxx
			if r.Offset()+2 > maxOffset {
				return ""
			}
			char2 := r.Read()
			char3 := r.Read()
			if (char2&0xC0) != 0x80 || (char3&0xC0) != 0x80 {
				return ""
			}
			charArray[charArrayOffset] = rune(((c & 0x0F) << 12) | ((char2 & 0x3F) << 6) | (char3 & 0x3F))
			charArrayOffset++
		default:
			// malformed input
			return ""
		}
	}

	return string(charArray[:charArrayOffset])
}
