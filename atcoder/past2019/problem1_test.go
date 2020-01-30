package main

import "testing"

func TestRun1(t *testing.T) {
	run1([]byte("123"))
	run1([]byte("1x3"))
	run1([]byte("000"))
	run1([]byte("012"))
	run1([]byte("002"))
	run1([]byte("999"))
	run1([]byte("-1"))
	run1([]byte("aaa"))
}
