// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package prefix

import (
	"fmt"
	"testing"
)

func TestTrieRoot_String(t *testing.T) {
	n := NewTrie()
	n.Insert("foo", "foo")
	n.Insert("foooo", "foooo")
	n.Insert("foooeee", "foooeee")
	fmt.Println(n.Format())
	fmt.Println(n.FindPrefix("foooo", 1))
	t.Fatal("zz")
}
