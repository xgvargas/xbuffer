package buffer

import (
	"fmt"
	"testing"
)

func TestWriteLittleEndian(t *testing.T) {
	b := NewBuffer(1+2+4+8, LittleEndian)

	b.WriteI8(0, 1)
	b.WriteI16(1, 0x203)
	b.WriteI32(3, 0x4050607)
	b.WriteI64(7, 0x8090a0b0c0d0e0f)

	res := b.GetSlice()

	expect := []byte{1, 3, 2, 7, 6, 5, 4, 15, 14, 13, 12, 11, 10, 9, 8}

	if len(res) != len(expect) {
		t.Errorf("Wrong size")
	}

	for i, v := range res {
		if v != expect[i] {
			t.Errorf("Result mismatch [%d] %d != %d", i, v, expect[i])
		}
	}
}

func TestWriteBigEndian(t *testing.T) {
	b := NewBuffer(1+2+4+8, BigEndian)

	b.WriteI8(0, 1)
	b.WriteI16(1, 0x203)
	b.WriteI32(3, 0x4050607)
	b.WriteI64(7, 0x8090a0b0c0d0e0f)

	res := b.GetSlice()

	expect := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	if len(res) != len(expect) {
		t.Errorf("Wrong size")
	}

	for i, v := range res {
		if v != expect[i] {
			t.Errorf("Result mismatch [%d] %d != %d", i, v, expect[i])
		}
	}
}

func ExampleNewBuffer_writeLlittleEndian() {
	b := NewBuffer(1+2+4+8+4+4, LittleEndian)

	b.WriteI8(0, 1)
	b.WriteI16(1, 0x203)
	b.WriteI32(3, 0x4050607)
	b.WriteI64(7, 0x8090a0b0c0d0e0f)
	b.WriteSlice(15, []byte{1, 2, 3, 4})
	b.WriteSlice(19, []byte("ABCD"))

	res := b.GetSlice()

	fmt.Println(res)

	// Output: [1 3 2 7 6 5 4 15 14 13 12 11 10 9 8 1 2 3 4 65 66 67 68]

}

func ExampleNewBuffer_writeBigEndian() {
	b := NewBuffer(1+2+4+8+4+4, BigEndian)

	b.WriteI8(0, 1)
	b.WriteI16(1, 0x203)
	b.WriteI32(3, 0x4050607)
	b.WriteI64(7, 0x8090a0b0c0d0e0f)
	b.WriteSlice(15, []byte{1, 2, 3, 4})
	b.WriteSlice(19, []byte("ABCD"))

	res := b.GetSlice()

	fmt.Println(res)

	// Output: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 1 2 3 4 65 66 67 68]
}

func ExampleNewBuffer_readLittleEndian() {
	b := NewBufferFrom([]byte{1, 3, 2, 7, 6, 5, 4, 15, 14, 13, 12, 11, 10, 9, 8, 65, 66, 67}, LittleEndian)

	r := b.ReadI8(0)
	s := b.ReadI16(1)
	t := b.ReadI32(3)
	u := b.ReadI64(7)
	v := string(b.ReadSlice(15, 3))

	fmt.Printf("0x%x, 0x%x, 0x%x, 0x%x, %v", r, s, t, u, v)

	// Output: 0x1, 0x203, 0x4050607, 0x8090a0b0c0d0e0f, ABC
}

func ExampleNewBuffer_readBigEndian() {
	b := NewBufferFrom([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 2, 3}, BigEndian)

	r := b.ReadI8(0)
	s := b.ReadI16(1)
	t := b.ReadI32(3)
	u := b.ReadI64(7)
	v := b.ReadSlice(15, 3)

	fmt.Printf("0x%x, 0x%x, 0x%x, 0x%x, %v", r, s, t, u, v)

	// Output: 0x1, 0x203, 0x4050607, 0x8090a0b0c0d0e0f, [1 2 3]
}
