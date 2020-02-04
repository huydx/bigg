package main

import (
	"bytes"
	"fmt"
	"sort"
)

type state int

func run6(input []byte) {
	var (
		insideString bool
		buf          bytes.Buffer
		dict         = make([]string, 0)
	)

	for _, b := range input {
		if upperCase(b) {
			if insideString {
				buf.Write([]byte{b})
				s := buf.String()
				buf.Reset()
				insideString = false
				dict = append(dict, s)
			} else {
				insideString = true
				buf.Write([]byte{b})
			}
		} else {
			buf.Write([]byte{b})
		}
	}

	sort.Slice(dict, func(i, j int) bool {
		return dict[i] < dict[j]
	})

	fmt.Println(dict)
}

func upperCase(b byte) bool {
	s := int(b) - 65
	return s >= 0 && s <= 25
}
