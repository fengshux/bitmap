package bitmap

import (
	"fmt"
)

type BitMap struct {
	store []uint64
	size  uint64
}

// new Bitmap to create a bitmap
// size is the length of total of bit
// return the pointer of Bitmap
func New(size uint64) *BitMap {

	length := size / 64

	if size%64 > 0 {
		length = length + 1
	}

	return &BitMap{make([]uint64, length), size}

}

// Set 1 on the offset bit.
// Panic when offset bigger then size of Bitmap
func (b *BitMap) Set(offset uint64) {

	b.checkOffset(offset)
	index, bit := indexAndOverbit(offset)
	count := uint64(1 << bit)

	b.store[index] = b.store[index] | count
}

// set 0 on the offset bit.
// Panic when offset bigger then size of Bitmap
func (b *BitMap) Clear(offset uint64) {

	b.checkOffset(offset)
	index, bit := indexAndOverbit(offset)
	count := ^uint64(1 << bit)
	b.store[index] = b.store[index] & count

}

// get value on the offset bit.
// Panic when offset bigger then size of Bitmap
func (b *BitMap) Get(offset uint64) int {

	b.checkOffset(offset)
	index, bit := indexAndOverbit(offset)
	count := uint64(1 << bit)
	val := b.store[index] & count

	if val == 0 {
		return 0
	}
	return 1
}

func (b *BitMap) checkOffset(offset uint64) {
	if offset > b.size {
		panic(fmt.Sprintf("BitMap overflow: offset %d bigger than size %d", offset, b.size))
	}
}

func indexAndOverbit(offset uint64) (index, bit uint64) {
	return offset / 64, offset % 64
}
