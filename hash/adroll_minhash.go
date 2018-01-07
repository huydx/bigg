package bigg

//an implementation using adroll paper: http://tech.adroll.com/media/hllminhash.pdf
//which fuse HLL and min-hash in the same data structure

//tree set implementation

type adrollMinHash struct {
	p         uint32 //num of bit taken to calculate HLL
	k         uint32 //num of hash functions
	precision uint32
}

func (a *adrollMinHash) add(item string) {

}

func (a *adrollMinHash) intersect(another *adrollMinHash) float32 {

}

func (a *adrollMinHash) union(another *adrollMinHash) float32 {

}

func (a *adrollMinHash) estimate() uint64 {

}
