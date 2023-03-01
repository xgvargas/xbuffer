/*

https://medium.com/learning-the-go-programming-language/encoding-data-with-the-go-binary-package-42c7c0eb3e73

*/

package xbuffer

type Order int

const (
	LittleEndian Order = iota + 1
	BigEndian
)

type Buffer struct {
	buf []byte
	dir Order
}

// creates a new Buffer starting `size` bytes long
func NewBuffer(size int, dir Order) Buffer {

	return Buffer{
		make([]byte, size),
		dir,
	}
}

// create a new Buffer and sets its content to a copy of `data`
func NewBufferFrom(data []byte, dir Order) Buffer {

	b := make([]byte, len(data))

	copy(b, data)

	return Buffer{
		b,
		dir,
	}
}

func (b *Buffer) fixLen(size int) {
	// if size > len(b.buf) {
	// 	// extra := make([]byte, size-len(b.buf))
	// 	// b.buf = append(b.buf, extra)
	// }
}

func (b *Buffer) GetSlice() []byte {
	return b.buf
}

func (b *Buffer) ReadSlice(offset, size int) []byte {
	tmp := make([]byte, size)
	copy(tmp, b.buf[offset:offset+size])
	return tmp
}

func (b *Buffer) ReadU8(offset int) uint8 {
	return b.buf[offset]
}

func (b *Buffer) ReadI8(offset int) int8 {
	return int8(b.buf[offset])
}

func (b *Buffer) ReadU16(offset int) uint16 {
	var d uint16 = 0
	if b.dir == LittleEndian {
		d |= uint16(b.buf[offset+0])
		d |= uint16(b.buf[offset+1]) << 8
	} else {
		d |= uint16(b.buf[offset+1])
		d |= uint16(b.buf[offset+0]) << 8
	}
	return d
}

func (b *Buffer) ReadI16(offset int) int16 {
	return int16(b.ReadU16(offset))
}

func (b *Buffer) ReadU32(offset int) uint32 {
	var d uint32 = 0
	if b.dir == LittleEndian {
		d |= uint32(b.buf[offset+0])
		d |= uint32(b.buf[offset+1]) << 8
		d |= uint32(b.buf[offset+2]) << 16
		d |= uint32(b.buf[offset+3]) << 24
	} else {
		d |= uint32(b.buf[offset+3])
		d |= uint32(b.buf[offset+2]) << 8
		d |= uint32(b.buf[offset+1]) << 16
		d |= uint32(b.buf[offset+0]) << 24
	}
	return d
}

func (b *Buffer) ReadI32(offset int) int32 {
	return int32(b.ReadU32(offset))
}

func (b *Buffer) ReadU64(offset int) uint64 {
	var d uint64 = 0
	if b.dir == LittleEndian {
		d |= uint64(b.buf[offset+0])
		d |= uint64(b.buf[offset+1]) << 8
		d |= uint64(b.buf[offset+2]) << 16
		d |= uint64(b.buf[offset+3]) << 24
		d |= uint64(b.buf[offset+4]) << 32
		d |= uint64(b.buf[offset+5]) << 40
		d |= uint64(b.buf[offset+6]) << 48
		d |= uint64(b.buf[offset+7]) << 56
	} else {
		d |= uint64(b.buf[offset+7])
		d |= uint64(b.buf[offset+6]) << 8
		d |= uint64(b.buf[offset+5]) << 16
		d |= uint64(b.buf[offset+4]) << 24
		d |= uint64(b.buf[offset+3]) << 32
		d |= uint64(b.buf[offset+2]) << 40
		d |= uint64(b.buf[offset+1]) << 48
		d |= uint64(b.buf[offset+0]) << 56
	}
	return d
}

func (b *Buffer) ReadI64(offset int) int64 {
	return int64(b.ReadU64(offset))
}

func (b *Buffer) WriteU8(offset int, val uint8) {
	b.fixLen(offset + 1)
	b.buf[offset] = val
}

func (b *Buffer) WriteI8(offset int, val int8) {
	b.buf[offset] = byte(val)
}

func (b *Buffer) WriteU16(offset int, val uint16) {
	b.fixLen(offset + 2)
	if b.dir == LittleEndian {
		b.buf[offset] = byte(val & 0xff)
		b.buf[offset+1] = byte(val >> 8)
	} else {
		b.buf[offset+1] = byte(val & 0xff)
		b.buf[offset] = byte(val >> 8)
	}
}

func (b *Buffer) WriteI16(offset int, val int16) {
	b.WriteU16(offset, uint16(val))
}

func (b *Buffer) WriteU32(offset int, val uint32) {
	b.fixLen(offset + 4)
	if b.dir == LittleEndian {
		b.buf[offset] = byte(val & 0xff)
		b.buf[offset+1] = byte((val >> 8) & 0xff)
		b.buf[offset+2] = byte((val >> 16) & 0xff)
		b.buf[offset+3] = byte(val >> 24)
	} else {
		b.buf[offset+3] = byte(val & 0xff)
		b.buf[offset+2] = byte((val >> 8) & 0xff)
		b.buf[offset+1] = byte((val >> 16) & 0xff)
		b.buf[offset] = byte(val >> 24)
	}

}

func (b *Buffer) WriteI32(offset int, val int32) {
	b.WriteU32(offset, uint32(val))
}

func (b *Buffer) WriteU64(offset int, val uint64) {
	b.fixLen(offset + 8)
	if b.dir == LittleEndian {
		b.buf[offset+0] = byte(val & 0xff)
		b.buf[offset+1] = byte((val >> 8) & 0xff)
		b.buf[offset+2] = byte((val >> 16) & 0xff)
		b.buf[offset+3] = byte((val >> 24) & 0xff)
		b.buf[offset+4] = byte((val >> 32) & 0xff)
		b.buf[offset+5] = byte((val >> 40) & 0xff)
		b.buf[offset+6] = byte((val >> 48) & 0xff)
		b.buf[offset+7] = byte(val >> 56)
	} else {
		b.buf[offset+7] = byte(val & 0xff)
		b.buf[offset+6] = byte((val >> 8) & 0xff)
		b.buf[offset+5] = byte((val >> 16) & 0xff)
		b.buf[offset+4] = byte((val >> 24) & 0xff)
		b.buf[offset+3] = byte((val >> 32) & 0xff)
		b.buf[offset+2] = byte((val >> 40) & 0xff)
		b.buf[offset+1] = byte((val >> 48) & 0xff)
		b.buf[offset+0] = byte(val >> 56)
	}

}

func (b *Buffer) WriteI64(offset int, val int64) {
	b.WriteU64(offset, uint64(val))
}

func (b *Buffer) WriteSlice(offset int, from []byte) {
	copy(b.buf[offset:], from)
}
