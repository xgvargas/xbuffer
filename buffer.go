/*

https://medium.com/learning-the-go-programming-language/encoding-data-with-the-go-binary-package-42c7c0eb3e73

*/

// Package xbuffer is a package to write and read data with a flexible format both in its size and endianess.
package xbuffer

import "fmt"

// Order define if buffer content is ordered as Little or Big endian
type Order int

const (
	LittleEndian Order = iota + 1
	BigEndian
)

// Xbuffer holds the internal buffer
type Xbuffer struct {
	buf []byte
	dir Order
}

// NewBuffer creates a new Xbuffer starting `size` bytes long
func NewBuffer(size int, dir Order) Xbuffer {

	return Xbuffer{
		make([]byte, size),
		dir,
	}
}

// NewBufferFrom create a new Xbuffer and sets its content to a copy of `data`
func NewBufferFrom(data []byte, dir Order) Xbuffer {

	b := make([]byte, len(data))

	copy(b, data)

	return Xbuffer{
		b,
		dir,
	}
}

func (b *Xbuffer) fixLen(last int) {

	if last > len(b.buf) {

		panic(fmt.Sprintf("Trying to write to address %d of %d bytes long xbuffer", last, len(b.buf)))

		// TODO optionally grow internal buffer if necessary
		// extra := make([]byte, last-len(b.buf))
		// b.buf = append(b.buf, extra)
	}
}

// GetSlice returns the internal buffer
func (b *Xbuffer) GetSlice() []byte {
	return b.buf
}

// ReadSlice returns a slice of internal buffer
func (b *Xbuffer) ReadSlice(offset, size int) []byte {
	tmp := make([]byte, size)
	copy(tmp, b.buf[offset:offset+size])
	return tmp
}

// ReadU8 returns unsigned 8 bits from offset in internal buffer
func (b *Xbuffer) ReadU8(offset int) uint8 {
	return b.buf[offset]
}

// ReadI8 returns signed 8 bits from offset in internal buffer
func (b *Xbuffer) ReadI8(offset int) int8 {
	return int8(b.buf[offset])
}

// ReadU16 returns unsigned 16 bits from offset in internal buffer
func (b *Xbuffer) ReadU16(offset int) uint16 {
	var d uint16
	if b.dir == LittleEndian {
		d |= uint16(b.buf[offset+0])
		d |= uint16(b.buf[offset+1]) << 8
	} else {
		d |= uint16(b.buf[offset+1])
		d |= uint16(b.buf[offset+0]) << 8
	}
	return d
}

// ReadI16 returns signed 16bits from offset in internal buffer
func (b *Xbuffer) ReadI16(offset int) int16 {
	return int16(b.ReadU16(offset))
}

// ReadU32 returns unsigned 32 bits from offset in internal buffer
func (b *Xbuffer) ReadU32(offset int) uint32 {
	var d uint32
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

// ReadI32 returns signed 32 bits from offset in internal buffer
func (b *Xbuffer) ReadI32(offset int) int32 {
	return int32(b.ReadU32(offset))
}

// ReadU64 returns unsigned 64 bits from offset in internal buffer
func (b *Xbuffer) ReadU64(offset int) uint64 {
	var d uint64
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

// ReadI64 returns signed 64 bits from offset in internal buffer
func (b *Xbuffer) ReadI64(offset int) int64 {
	return int64(b.ReadU64(offset))
}

// WriteU8 write unsigned 8 bits value to internal buffer at offset
func (b *Xbuffer) WriteU8(offset int, val uint8) {
	b.fixLen(offset + 1)
	b.buf[offset] = val
}

// WriteI8 write signed 8 bits value to internal buffer at offset
func (b *Xbuffer) WriteI8(offset int, val int8) {
	b.buf[offset] = byte(val)
}

// WriteU16 write unsigned 16 bits value to internal buffer at offset
func (b *Xbuffer) WriteU16(offset int, val uint16) {
	b.fixLen(offset + 2)
	if b.dir == LittleEndian {
		b.buf[offset] = byte(val & 0xff)
		b.buf[offset+1] = byte(val >> 8)
	} else {
		b.buf[offset+1] = byte(val & 0xff)
		b.buf[offset] = byte(val >> 8)
	}
}

// WriteI16 write signed 16 bits value to internal buffer at offset
func (b *Xbuffer) WriteI16(offset int, val int16) {
	b.WriteU16(offset, uint16(val))
}

// WriteU32 write unsigned 32 bits value to internal buffer at offset
func (b *Xbuffer) WriteU32(offset int, val uint32) {
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

// WriteI32 write signed 32 bits value to internal buffer at offset
func (b *Xbuffer) WriteI32(offset int, val int32) {
	b.WriteU32(offset, uint32(val))
}

// WriteU64 write unsigned 64 bits value to internal buffer at offset
func (b *Xbuffer) WriteU64(offset int, val uint64) {
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

// WriteI64 write signed 64 bits value to internal buffer at offset
func (b *Xbuffer) WriteI64(offset int, val int64) {
	b.WriteU64(offset, uint64(val))
}

// WriteSlice write a slice to internal buffer at offset
func (b *Xbuffer) WriteSlice(offset int, from []byte) {
	b.fixLen(offset + len(from))
	copy(b.buf[offset:], from)
}
