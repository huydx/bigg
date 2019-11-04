// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package main

import (
	"fmt"
	"os"
	"time"
)

const (
	pageSize = 32 * 100
)

type WAL1 struct {
	segment *os.File
	buf     [pageSize]byte
	bufIdx  int
}

func (w *WAL1) run() {
	for {
		time.Sleep(time.Second * 2)
		s := time.Now()
		if err := w.segment.Sync(); err != nil {
			panic(err)
		}
		fmt.Printf("fsync cost %f ms\n", float64(time.Since(s).Nanoseconds())/1000000)
	}
}

func (w *WAL1) log(records ...[]byte) error {
	for _, rec := range records {
		if w.bufIdx+len(rec) > pageSize {
			_, err := w.segment.Write(rec)
			if err != nil {
				return err
			}
			for i := range w.buf {
				w.buf[i] = 0
			}
			w.bufIdx = 0
		}
		copy(w.buf[w.bufIdx:], rec)
		w.bufIdx += len(rec)
	}
	return nil
}

func main() {
	s, err := os.OpenFile("/tmp/segment.1", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	w := WAL1{
		segment: s,
	}

	bs := [32]byte{}
	for i, _ := range bs {
		bs[i] = 'A'
	}
	go w.run()
	for {
		st := time.Now()
		for i := 0; i < 100000; i++ {
			err = w.log(bs[0:30])
			if err != nil {
				panic(err)
			}
		}
		fmt.Printf("write 100000 records cost %f ms\n", float64(time.Since(st).Nanoseconds())/1000000)
	}
}
