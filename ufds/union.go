package udfs

// array p, rank, setSize
// int numSets

type Unionfind struct {
	p       []int
	rank    []int
	setSize []int
	numSets int
}

func NewUnionfind(N int) *Unionfind {
	p := make([]int, N)
	rank := make([]int, N)
	setSize := make([]int, N)

	for i := 0; i < N; i++ {
		p[i] = i
		rank[i] = 0
		setSize[i] = 1
	}

	return &Unionfind{
		p:       p,
		rank:    rank,
		setSize: setSize,
		numSets: N,
	}
}

func (u *Unionfind) findSet(i int) int {
	if u.p[i] == i {
		return i
	} else {
		ret := u.findSet(u.p[i])
		u.p[i] = ret
		return ret
	}
}

func (u *Unionfind) isSameSet(i int, j int) bool {
	return u.findSet(i) == u.findSet(j)
}

func (u *Unionfind) unionSet(i int, j int) {
	if !u.isSameSet(i, j) {
		u.numSets--
		x := u.findSet(i)
		y := u.findSet(j)

		if u.rank[x] > u.rank[y] {
			u.p[y] = x
			u.setSize[x] = u.setSize[x] + u.setSize[y]
		} else {
			u.p[x] = y
			u.setSize[y] = u.setSize[y] + u.setSize[x]

			if u.rank[x] == u.rank[y] {
				u.rank[y] = u.rank[y] + 1
			}
		}
	}
}

func (u *Unionfind) numDisjointSets() int {
	return u.numSets
}

func (u *Unionfind) sizeOfSet(i int) int {
	return u.setSize[u.findSet(i)]
}
