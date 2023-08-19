package rand

const (
	n = 624
	w = 32
)

type MersenneTwister struct {
	state [n]uint32
	index int
}

func NewMersenneTwister(seed uint32) *MersenneTwister {
	const f = 1812433253

	mt := MersenneTwister{}

	mt.state[0] = seed
	for i := 1; i < len(mt.state); i++ {
		x := mt.state[i-1]
		mt.state[i] = f*(x^(x>>(w-2))) + uint32(i)
	}

	mt.twist()
	return &mt
}

func MersenneTwisterFromState(state [n]uint32) *MersenneTwister {
	mt := MersenneTwister{state: state}
	mt.index = n
	return &mt
}

func (mt *MersenneTwister) Next() uint32 {
	if mt.index >= n {
		mt.twist()
	}
	r := Temper(mt.state[mt.index])
	mt.index++
	return r
}

func Temper(x uint32) uint32 {
	const (
		u = 11
		d = 0xFFFFFFFF
		s = 7
		b = 0x9D2C5680
		t = 15
		c = 0xEFC60000
		l = 18
	)

	y := x ^ ((x >> u) & d)
	y ^= ((y << s) & b)
	y ^= ((y << t) & c)
	return y ^ (y >> l)
}

func (mt *MersenneTwister) twist() {
	const (
		a         = 0x9908b0df
		m         = 397
		upperMask = 0x80000000
		lowerMask = 0x7fffffff
	)

	for i := 0; i < len(mt.state); i++ {
		x := (mt.state[i] & upperMask) + (mt.state[(i+1)%n] & lowerMask)
		xA := x >> 1
		if x%2 != 0 {
			xA ^= a
		}
		mt.state[i] = mt.state[(i+m)%n] ^ xA
	}
	mt.index = 0
}
