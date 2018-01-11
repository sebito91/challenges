// Package circular provides a circular-buffer implementation with overflow-checked writes
// and unconditional writes
package circular

import (
	"fmt"
	"log"
)

const testVersion = 4

// Buffer is a custom type for this project
type Buffer struct {
	data  []byte
	size  int
	write int
	read  int
}

// NewBuffer instantiates a new instance of our Buffer struct
func NewBuffer(size int) *Buffer {
	return &Buffer{size: size}
}

// ReadByte returns the next byte in the stack
func (b *Buffer) ReadByte() (out byte, err error) {
	if b.size < 1 || len(b.data) < 1 {
		return 0, fmt.Errorf("empty data slice")
	}

	if b.read >= len(b.data) {
		b.read = 0
	}
	out = b.data[b.read]
	b.read++
	b.size--

	return out, err
}

// checkInit is a quick helper function to check if we're zero
func (b *Buffer) checkInit() {
	if len(b.data) < 1 {
		b.data = make([]byte, b.size)
		b.size = 0
	}
}

// WriteByte appends the next byte onto the stack
func (b *Buffer) WriteByte(c byte) error {
	b.checkInit()
	if b.size >= len(b.data) {
		return fmt.Errorf("full buffer")
	}

	b.data[b.write] = c
	b.write++
	b.size++

	if b.write == len(b.data) {
		b.write = 0 // wrap
	}

	return nil
}

// Overwrite blindly updates the current pointer
func (b *Buffer) Overwrite(c byte) {
	b.checkInit()
	if b.size < len(b.data) {
		if err := b.WriteByte(c); err != nil {
			log.Printf("%v", err)
		}
	} else {
		b.data[b.read] = c
		b.read++
		if b.read == len(b.data) {
			b.read = 0
		}
	}
}

// Reset will return the data to nothing!
func (b *Buffer) Reset() {
	b.size = len(b.data)
	b.data = b.data[:0]
}
