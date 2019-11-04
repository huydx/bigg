

package main

import (
	"io"
	"bytes"
)

type RingBuffer struct {
	io.Writer
	io.Reader
	length int
	tail   int
	head   int
	buff   bytes.Buffer
}

func NewRingBuffer(len int) {
	r := &RingBuffer{
		length: len,
		buff:   bytes.NewBuffer([]byte{}),
	}
	r.buff.Grow(len)
	return r
}

func (r *RingBuffer) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (r *RingBuffer) Read(p []byte) (n int, err error) {
	return 0, nil
}
