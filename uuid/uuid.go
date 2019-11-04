package uuid

type UUID [16]byte // 128bit

// layout for RFC4122
//0                   1                   2                   3
//0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//|                          time_low                             |
//+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//|       time_mid                |         time_hi_and_version   |
//+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//|clk_seq_hi_res |  clk_seq_low  |         node (0-1)            |
//+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//|                         node (2-5)                            |
//+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

type Version int
const (
	V1 = iota // RFC4122
	V2
	V3
	V4
)

func (u *UUID) String() string {
	return ""
}

func (u *UUID) Version() byte {
	return u[6] >> 4
}

func NewUUID(ver Version) *UUID {
	switch ver {
	case V1:
		return genRFC4122UUID()
	case V2, V3, V4:
		return nil
	default:
		panic("unrecognize uuid version")
	}
}

func genRFC4122UUID() *UUID {
	return nil
}

func getClockSequence() (uint64, uint16, error) {

}
