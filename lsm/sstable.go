// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package lsm

type index struct {
	key    string
	offset int64
}

type sstable struct {
	indices []index
}
