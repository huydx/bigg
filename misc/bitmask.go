package main

import "fmt"

type bitmask struct {
	mask int64
}

func (b *bitmask) set(j int) {
	b.mask |= 1 << uint(j)
}

func (b *bitmask) check(j int) bool {
	return (b.mask ^ 1 << uint(j)) == 0
}

func (b *bitmask) clear(j int) {
	b.mask &= ^(1 << uint(j))
}

func (b *bitmask) toggle(j int) {
	b.mask ^= 1 << uint(j)
}

func (b *bitmask) lsb() int {
	return int(b.mask & -b.mask)
}

func main() {
	s := 39
	ss := 36
	fmt.Println((s+1) & s)
	fmt.Println((ss-1) | s)

	b := bitmask{}
	b.set(15)
	fmt.Println(b.check(15))
	b.clear(15)
	fmt.Println(b.check(15))
	b.set(15)
	fmt.Println(b.check(15))
	fmt.Println(b.check(1))
	b.toggle(1)
	fmt.Println(b.check(1))
}
