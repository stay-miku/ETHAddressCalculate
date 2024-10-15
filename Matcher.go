package main

type Matcher interface {
	MatchPrefix([]byte) bool
	MatchSuffix([]byte) bool
	MatchAll([]byte) bool
	MatchOne([]byte) bool
}

// ETHMatcher prefix and suffix length must over or equal 2
type ETHMatcher struct {
	prefixLen int
	suffixLen int
}

type TronMatcher struct {
	prefixLen int
	suffixLen int
}

// MatchPrefix is unsafe, you need to ensure data length is greater than prefixLen
func (m *ETHMatcher) MatchPrefix(data []byte) bool {
	if !(data[0]>>4 == data[0]&0x0F) {
		return false
	}
	var match byte = data[0]
	byteLen := m.prefixLen % 2
	for i := 1; i < m.prefixLen/2; i++ {
		if match^data[i] != 0 {
			return false
		}
	}
	if byteLen != 0 {
		return (data[m.prefixLen/2] >> 4) == (match & 0x0F)
	}
	return true
	//data = []byte(hex.EncodeToString(data))
	//var match byte = data[0]
	//for i := 1; i < m.prefixLen; i++ {
	//	if match^data[i] != 0 {
	//		return false
	//	}
	//}
	//return true
}

func (m *ETHMatcher) MatchSuffix(data []byte) bool {
	length := len(data)
	if !(data[length-1]>>4 == data[length-1]&0x0F) {
		return false
	}

	var match byte = data[length-1]
	byteLen := m.suffixLen % 2
	for i := length - 2; i > length-m.suffixLen/2-1; i-- {
		if match^data[i] != 0 {
			return false
		}
	}
	if byteLen != 0 {
		return (data[length-m.suffixLen/2-1] & 0x0F) == (match >> 4)
	}
	return true

	//data = []byte(hex.EncodeToString(data))
	//length := len(data)
	//match := data[length-1]
	//for i := length - 2; i > length-m.suffixLen-1; i-- {
	//	if match^data[i] != 0 {
	//		return false
	//	}
	//}
	//return true
}

func (m *ETHMatcher) MatchAll(data []byte) bool {
	return m.MatchPrefix(data) && m.MatchSuffix(data)
}

func (m *ETHMatcher) MatchOne(data []byte) bool {
	return m.MatchPrefix(data) || m.MatchSuffix(data)
}

// MatchPrefix Tron Address is start with T
func (m *TronMatcher) MatchPrefix(data []byte) bool {
	var match byte = 'T'
	for i := 1; i < m.prefixLen; i++ {
		if match^data[i] != 0 {
			return false
		}
	}
	return true
}

func (m *TronMatcher) MatchSuffix(data []byte) bool {
	length := len(data)
	match := data[length-1]
	for i := length - 2; i > length-m.suffixLen-1; i-- {
		if match^data[i] != 0 {
			return false
		}
	}
	return true
}

func (m *TronMatcher) MatchAll(data []byte) bool {
	return m.MatchPrefix(data) && m.MatchSuffix(data)
}

func (m *TronMatcher) MatchOne(data []byte) bool {
	return m.MatchPrefix(data) || m.MatchSuffix(data)
}

func NewETHMatcher(prefixLen, suffixLen int) Matcher {
	return &ETHMatcher{
		prefixLen: prefixLen,
		suffixLen: suffixLen,
	}
}

func NewTronMatcher(prefixLen, suffixLen int) Matcher {
	return &TronMatcher{
		prefixLen: prefixLen,
		suffixLen: suffixLen,
	}
}
