

package dtab

import "fmt"

type NameTreeType interface{}

type Dentry struct {
	cursor int
	origin string
	Prefix string
	Suffix NameTree
}

func (dentry *Dentry) String() string {
	return fmt.Sprintf("%s => %s", dentry.Prefix, dentry.Suffix)
}

func (d *Dentry) Parse(input string) *Dentry {
	prefix := d.parsePrefix()
}

func (dentry *Dentry) parsePrefix() string {

}

func (dentry *Dentry) parseWhiteSpace() {
	for {
		if dentry.origin[dentry.cursor] == ' ' {
			dentry.cursor++
		} else {
			break
		}
	}
}

func (dentry *Dentry) parseCharacter(c byte) error {
	if dentry.origin[dentry.cursor] != c {
		return fmt.Errorf("expect %c found %c", c, dentry.origin[dentry.cursor])
	}
	dentry.cursor++
	return nil
}

type NameTree struct {
}

func (n *NameTree) Parse(input string) {

}

func (n1 *NameTree) And(n2 *NameTree) *NameTree {

}
