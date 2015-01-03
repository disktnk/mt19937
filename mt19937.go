package mt19937

/* http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/VERSIONS/C-LANG/mt19937-64.c */
const (
	N                 = 312
	M                 = 156
	MATRIX_A   uint64 = 0xB5026F5AA96619E9
	UPPER_MASK uint64 = 0xFFFFFFFF80000000
	LOWER_MASK uint64 = 0x7FFFFFFF

	NO_SEED = N + 1
)

type MT19937_64 struct {
	state [NO_SEED]uint64
	index int
}

func New() *MT19937_64 {
	return &MT19937_64{
		index: NO_SEED,
	}
}

func (m *MT19937_64) Seed(seed int64) {
	m.state[0] = uint64(seed)
	for i := uint64(1); i < N; i++ {
		m.state[i] = 6364136223846793005*(m.state[i-1]^m.state[i-1]>>62) + i
	}
	m.index = N
}

/* initialize by an array with array-length
   init_key is the array for initializing keys */
func (m *MT19937_64) SeedByArray(key []uint64) {
	m.Seed(19650218)
	i := 1
	j := 0
	k := len(key)
	if N > k {
		k = N
	}
	for k > 0 {
		m.state[i] = (m.state[i] ^ ((m.state[i-1] ^ (m.state[i-1] >> 62)) * 3935559000370003845) +
			key[j] + uint64(j))
		i++
		j++
		if i >= N {
			m.state[0] = m.state[N-1]
			i = 1
		}
		if j >= len(key) {
			j = 0
		}
		k--
	}
	for j = 0; j < N-1; j++ {
		m.state[i] = m.state[i] ^ ((m.state[i-1] ^ (m.state[i-1] >> 62)) * 2862933555777941757) -
			uint64(i)
		i++
		if i >= N {
			m.state[0] = m.state[N-1]
			i = 1
		}
	}
	m.state[0] = 1 << 63
}

/* generates a (pseudo-)random number on [0, 2^64-1]-interval */
func (m *MT19937_64) Uint64() uint64 {
	var i int
	var x uint64
	mag01 := []uint64{0, MATRIX_A}
	if m.index >= N {
		if m.index == NO_SEED {
			m.Seed(int64(5489))
		}
		for i = 0; i < N-M; i++ {
			x = (m.state[i] & UPPER_MASK) | (m.state[i+1] & LOWER_MASK)
			m.state[i] = m.state[i+M] ^ (x >> 1) ^ mag01[int(x&uint64(1))]
		}
		for ; i < N-1; i++ {
			x = (m.state[i] & UPPER_MASK) | (m.state[i+1] & LOWER_MASK)
			m.state[i] = m.state[i+M-N] ^ (x >> 1) ^ mag01[int(x&uint64(1))]
		}
		x = (m.state[N-1] & UPPER_MASK) | (m.state[0] & LOWER_MASK)
		m.state[N-1] = m.state[M-1] ^ (x >> 1) ^ mag01[int(x&uint64(1))]
		m.index = 0
	}
	x = m.state[m.index]
	m.index++
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)
	return x
}

/* generates a (pseudo-)random number on [0, 2^63-1]-interval */
func (m *MT19937_64) Int63() int64 {
	return int64(m.Uint64() >> 1)
}

/* generates a (pseudo-)random number on [0,1]-real-interval */
func (m *MT19937_64) Float63_1() float64 {
	return float64(m.Uint64()>>11) / 9007199254740991.0
}

/* generates a (pseudo-)random number on [0,1)-real-interval */
func (m *MT19937_64) Float63() float64 {
	return float64(m.Uint64()>>11) / 9007199254740992.0
}

/* generates a (pseudo-)random number on (0,1)-real-interval */
func (m *MT19937_64) Float63_3() float64 {
	return (float64(m.Uint64()>>12) + 0.5) / 4503599627370496.0
}
