package filter

import "kvdb/pkg/utils"

type BloomFilter struct {
	bitsPerKey int
	hashKeys   []uint32
}

func (b *BloomFilter) Add(key []byte) {
	b.hashKeys = append(b.hashKeys, utils.Hash(key, 0xa1b2c3d4))
}

func (b *BloomFilter) Hash() []byte {

	n := len(b.hashKeys)
	k := uint8(b.bitsPerKey * 69 / (100 * n))

	if k < 1 {
		k = 1
	} else if k > 30 {
		k = 30
	}
	// 布隆过滤器bit数组长度
	nBits := uint32(n * b.bitsPerKey)
	nBytes := (nBits + 7) / 8
	nBits = nBytes * 8

	if nBits < 64 {
		nBits = 64
	}

	dest := make([]byte, nBytes+1)
	dest[nBytes] = k

	// hash1(key)+i*hash2(key)
	for _, h := range b.hashKeys {
		delta := (h >> 17) | (h << 15)
		for i := uint8(0); i < k; i++ {
			bitpos := h % nBits
			dest[bitpos/8] |= 1 << (bitpos % 8)
			h += delta
		}
	}

	// b.hashKeys = b.hashKeys[:0]
	return dest
}

func (b *BloomFilter) Size() int {
	n := len(b.hashKeys)
	k := uint8(b.bitsPerKey * 69 / (100 * n))

	if k < 1 {
		k = 1
	} else if k > 30 {
		k = 30
	}
	return int(n * b.bitsPerKey)
}

func (b *BloomFilter) KeyLen() int {
	return len(b.hashKeys)
}

func (b *BloomFilter) Reset() {
	b.hashKeys = b.hashKeys[:0]
}

func NewBloomFilter(bitsPerKey int) *BloomFilter {

	return &BloomFilter{
		bitsPerKey: bitsPerKey,
	}
}

func Contains(filter, key []byte) bool {

	nBytes := len(filter) - 1

	if nBytes < 1 {
		return false
	}

	nBits := uint32(nBytes * 8)
	k := filter[nBytes]

	if k > 30 {
		return true
	}

	h := utils.Hash(key, 0xa1b2c3d4)

	delta := (h >> 17) | (h << 15)
	for i := uint8(0); i < k; i++ {
		bitpos := h % nBits

		if (uint32(filter[bitpos/8]) & (1 << (bitpos % 8))) == 0 {
			return false
		}
		h += delta
	}
	return true
}