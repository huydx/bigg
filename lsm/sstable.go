

package lsm

type index struct {
	key    string
	offset int64
}

type sstable struct {
	indices []index
}
