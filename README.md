# xbuffer

A simple buffer for Go.

```go
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
```
